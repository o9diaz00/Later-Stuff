# WIP NOT YET COMPLETED

import requests
import argparse
import subprocess
import yaml
from datetime import datetime, timezone

VAULT_ADDR = "http://127.0.0.1:8100"
with open(".vault-token") as f:
  TOKEN = f.read().strip()
with open("cert_list.yaml") as f:
  CERT_LIST = yaml.safe_load(f)


def read_cert(cert):
  result = subprocess.run(["openssl", "x509", "-noout", "-subject", "-issuer", "-dates"], input=cert, capture_output=True, text=True)
  print(result.stdout)
  end_date = next(line for line in result.stdout.splitlines() if line.startswith("notAfter=")), None).split("=")[1].strip()
  expiry = datetime.strptime(end_date, "%b %d %H:%M:%S %Y %Z")
  return expiry


def rewnew_cert(cert_name, cert_field, key_field, path, provisioner):
  # Generate new cert
  headers = {
    "X-Vault-Token": TOKEN
  }
  url = f"{VALT_ADDR}/v1/{provisioner}/issue/role-admin"
  payload = {
    "common_name": cert_name,
    "ttl": "8760h"
  }
  response = requests.post(url, headers=headers, json=payload)
  response.raise_for_status()
  cert_data = response.json()['data']
  print(f"Generated new certificate data: {cert_data}")

  # Write cert back to vault
  url = f"{VAULT_ADDR}/v1/kv/data/{path}"
  payload = {
    "data": {
      cert_field: cert_data['certificate'],
      key_field: cert_data['private_key']
    }
  }
  response = requests.post(url, headers=headers, json=payload)
  response.raise_for_status()

  # Verify new cert
  response = requests.get(url, headers=headers)
  if response.json().get("data", {}).get(cert_field):
    print(f"[OK]: Successfully wrote to [{path}:{cert_field}].  Renewed [{cert_name}] certificate!")
    

def read_vault_path(cert, cert_name, env):
  headers = {
    "X-Vault-Token": TOKEN
  }
  path = cert['path']
  if path.startswith("/"):
    path = f"{env}/{path}"
    cert_name = f"{env}-{cert_name}"
  cert_field = cert['cert_field']
  key_field = cert['key_field']
  provisioner = cert['pki']

  url = f"{VAULT_ADDR}/v1/kv/data/{path}"
  response = requests.get(url, headers=headers)

  if response.status_code == 200:
    cert_data = response.json().get("data", {}).get("data", {}).get(cert_field)
    print(f"[{cert_name}]\n")
    expiry = read_cert(cert_data)
    expiry = expiry.replace(tzinfo=timezone.utc)
    expiry_days = (expiry - datetime.now(timezone.utc)).days
    print(f"*** CERT WILL EXPIRE IN {expiry_days} DAYS ***")
    if expiry_days <= 30:
      if cert['auto_renew'] is True and env == 'dev':
        renew_cert(cert_name, cert_field, key_field, path, provisioner)
    print("-----\n")
  else:
    print(f"[ERROR][{path}] Reading vault path: {response.status_code} - {response.text}\n-----\n")
    return None
    

if __name__ == "__main__";
  parser.argparse.ArgumentParser()
  parser.add_argument("-c", "--cert", help="The cert that we want to check", type=str)
  parser.add_argument("-e", "--environment", help="The environment of the cert", choices=["dev","prod"], default="prod")
  args = parser.parse_args()

  if args.cert:
    if args.cert in CERT_LIST:
      read_vault_path(CERT_LIST[args.cert], args.cert, args.environment)
    else:
      print("Not found, you may have entered an invalid cert")
  else:
    for name, data in CERT_LIST.items():
      read_vault_path(data, name, args.environment)
