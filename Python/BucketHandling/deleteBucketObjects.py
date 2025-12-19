import argparse
import boto3
import datetime
import re
from concurrent.futures import ThreadPoolExecutor, as_completed
import logging

args = None
SESSION = None
S3 = None

def get_accounts(bucket):
  account_list = []
  pages = S3.get_paginator('list_objects_v2').paginate(Bucket=bucket, Prefix="data/", Delimiter='/')
  for page in pages:
    for content in page.get('CommonPrefixes', []):
      account_list.append(content["Prefix"])

  return account_list

def get_delete_days(prefix, retention, bucket):
  delete_list = []
  now_utc = datetime.datetime.now(datetime.timezone.utc)
  dt = (now_utc - datetime.timedelta(days=retention)).date()
  days_regex = re.compile(r"day=(\d{4}-\d{2}-\d{2})/")
  pages = S3.get_paginator('list_objects_v2').paginate(Bucket=bucket, Prefix=prefix, Delimiter='/')
  for page in pages:
    for content in page.get('CommonPrefixes', []):
      match = days_regex.search(content["Prefix"])
      if not match:
        continue
      day = datetime.datetime.strptime(match.group(1), "%Y-%m-%d").date()
      if day <= dt:
        delete_list.append(content["Prefix"])

  return delete_list

def delete_objects(prefix, bucket):
  deleted = 0
  pages = S3.get_paginator('list_objects_v2').paginate(Bucket=bucket, Prefix=prefix)
  try:
    for page in pages:
      content = page.get("Contents", [])
      if not content:
        continue
      batch = [{"Key": c["Key"]} for c in content]
      for i in range(0, len(batch), 1000):
        if args.dry_run:
          logging.info(f"[DRY-RUN] Deleted {len(batch[i:i+1000])} objects from {prefix}")
        else:
          S3.delete_objects(Bucket=bucket, Delete={"Objects": batch[i:i+1000], "Quiet": True})
          logging.info(f"Deleted {len(batch[i:i+1000])} objects from {prefix}")
        deleted += len(batch[i:i+1000])
  except Exception as e:
    logging.error(f"Error deleting {prefix}: {e}")

  return prefix, deleted

def main():
  global args, SESSION, S3
  parser = argparse.ArgumentParser()
  parser.add_argument("-b", "--bucket", help="The bucket in which we want to check", type=str, required=True)
  parser.add_argument("-a", "--account", help="The account in which we want to check", type=str, required=True)
  parser.add_argument("-r", "--region", help="The region where the bucket exists", type=str, default="us-iad-12")
  parser.add_argument("-R", "--retention", help="The amount of days you want to delete up to", type=int, default=3)
  parser.add_argument("-w", "--workers", help="The number of concurrent workers", type=int, default=16)
  parser.add_argument("--dry-run", help="Prints what would be deleted without performing the action", action="store_true")
  args = parser.parse_args()

  SESSION = boto3.Session(profile_name=args.account)
  S3 = SESSION.client("s3", endpoint_url=f"https://{args.region}.linodeobjects.com")
  delete_prefix = []

  logging.basicConfig(
      level=logging.INFO,
      format="[%(asctime)s] [%(levelname)s] %(message)s",
      handlers=[
          logging.StreamHandler(),
          logging.FileHandler(f"_deleted.{args.bucket}_{args.region}.log")
      ]
  )

  logging.info(f"Cleaning up bucket [{args.bucket}]...\n - Account: {args.account}\n - Retention: {args.retention}\n - Region: {args.region}\n - Log File: _deleted.{args.bucket}_{args.region}.log")
  accounts = get_accounts(args.bucket)
  for a in accounts:
    delete_prefix.extend(get_delete_days(a, args.retention, args.bucket))

  with ThreadPoolExecutor(max_workers=args.workers) as executor:
    futures = [executor.submit(delete_objects, prefix, args.bucket) for prefix in delete_prefix]
    for f in as_completed(futures):
      prefix, count = f.result()


if __name__ == "__main__":
  main()
