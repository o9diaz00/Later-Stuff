import boto3
import argparse
import math
import datetime

def convert_bytes(size):
    t = ["B", "K", "M", "G", "T", "P"]
    n = 0
    
    while math.floor(size) > 1000:
        size /= 1024
        n += 1
        
    return size, t[n]

def get_bucket_data(bucket, path):
    data = []
    s3 = boto3.client('s3')
    pages = s3.get_paginator('list_objects_v2').paginate(Bucket=bucket, Prefix=path) if path else s3.get_paginator('list_objects_v2').paginate(Bucket=bucket)
    
    for page in pages:
        for content in page['Contents']:
            data.append(content)
            
    return data
    
def get_bucket_last_file(data):
    sorted_obj = sorted(data, key=lambda x: x['LastModified'], reverse=True)
    return sorted_obj[0]['Key'], str(sorted_obj[0]['LastModified'])
    
def get_bucket_object_count(data):
    return len(data)
    
def get_bucket_total_size(data):
    size = 0
    
    for n in data:
        size += n['Size']
        
    return size
    
def print_output(data, args):
    bucket_str = f"{args.bucket}/{args.path}" if args.path else f"{args.bucket}"
    print(f"[{bucket_str}]")
    if args.count:
        count = get_bucket_object_count(data)
        print(f" - Total Objects: {count}")
    if args.size:
        size = convert_bytes(get_bucket_total_size(data))
        print(f" - Total Size: {size[0]:.2f}{size[1]")
    if args.last:
        last = get_bucket_last_file(data)
        print(f" - Last Modified Object: {last[0]}\n - Last Modified Time: {last[1]}")
        
def print_alert(data, args):
    status = 0
    bucket_str = f"{args.bucket}/{args.path}" if args.path else f"{args.bucket}"
    
    if args.alert == "size":
        cur_size = get_bucket_total_size(data)
        size_limit = 1024
        if cur_size > size_limit:
            status = 1
            size = convert_bytes(cur_size)
            print(f"CRITICAL - [{bucket_str}] size is {size[0]:.2f}{size[1]!")       
    if args.alert == "count":
        count_limit = 2
        cur_count = get_bucket_object_count(data)
        if cur_count > count_limit:
            status = 1
            print(f"CRITICAL - [{bucket_str}] has {cur_count} objects!")
    if args.alert == "last":
        cur_hour = datetime.datetime.now().hour
        args.path += str(cur_hour)
        last_upload = get_bucket_last_file(data)
        if last_upload[1] < str(datetime.datetime.utcnow() - datetime.timedelta(hours=2)):
            print(f"CRITICAL - [{bucket_str}]'s last file was {last_upload[1]} which is more than 2 hours ago!")
            status = 1
    
    return status
    
if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-b", "--bucket", help="The bucket in which we want to check", type=str, required=True)
    parser.add_argument("-p", "--path", help="The subdirectory withi the bucket we want to check", type=str, required=False)
    parser.add_argument("-s", "--size", help="Output total size of bucket", action='store_true', required=False)
    parser.add_argument("-c", "--count", help="Output count of objects in bucket", action='store_true', required=False)
    parser.add_argument("-l", "--last", help="Output the last object uploaded to bucket", action='store_true', required=False)
    parser.add_argument("-a", "--all", help="Turn this on to check all accounts within the bucket", action='store_true', required=False)
    parser.add_argument("-A", "--alert", help="Choose metric to check and alert on ['size', 'count', 'last']", choices=["size", "count", "last"], required=False)
    args = parser.parse_args()
    EXIT_STATUS = 0
    
    if not args.all:
        DATA = get_bucket_data(args.bucket, args.path)
        if not args.alert:
            print_output(DATA, args)
        else:
            EXIT_STATUS = print_alert(DATA, args)
    else:
        ACCOUNTS = ["aaa","bbb","ccc"]
        for account in ACCOUNTS:
            args.path = "123/456=" + account + "/"
            DATA = get_bucket_data(args.bucket, args.path)
            if not args.alert:
                print_output(DATA, args)
            else:
                last_upload_date = get_bucket_last_file(DATA)
                args.path += "days=" + last_upload_date[1].split(' ')[0] + "/"
                DATA = get_bucket_data(args.bucket, args.path)
                EXIT_STATUS = print_alert(DATA, args)
                
    exit(EXIT_STATUS)
