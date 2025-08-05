data "ruvds_ssh_list" "keys" {
}

output "keys" {
  value = data.ruvds_ssh_list.keys
}

# Get list of all SSH keys deployed to RuVDS.
# Output:
#  + keys               = {
#      + ssh_keys = [
#          + {
#              + md5_fingerprint    = "MD5:cc:01:02:03:a9:b3:af:1d:8b:47:0c:0b:0a:03:02:01"
#              + name               = "MacBook"
#              + public_key         = "ssh-ed25519 AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
#              + sha256_fingerprint = "SHA256:KRU45kLJ3HgtxJzdh3JhsJMNAgtUneH/JFkJ6BAgs72"
#              + ssh_key_id         = "bc123456-0123-456ef-efef-00155d0093043"
#            },
#        ]
#    }
