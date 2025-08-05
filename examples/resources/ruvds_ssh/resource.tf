# Importing a SSH key by name
resource "ruvds_ssh" "my_key" {
  name = "MacBook"
}

# > terraform import ruvds_ssh.my_key "MacBook"

# Output:
#ruvds_ssh.my_key: Importing from ID "MacBook"...
#ruvds_ssh.my_key: Import prepared!
#  Prepared ruvds_ssh for import
#ruvds_ssh.my_key: Refreshing state... [name=MacBook]
#Import successful!

# > terraform plan

# Output:
#  # ruvds_ssh.my_key will be updated in-place
#  ~ resource "ruvds_ssh" "my_key" {
#      ~ md5_fingerprint    = "MD5:aa:aa:aa:aa:aa:aa:aa:aa:aa:aa:aa:aa:14:b1:a2:ce" -> (known after apply)
#        name               = "MacBook"
#      - public_key         = "ssh-ed25519 AAAAC**********nCC92urGnbwZ6iU/GuWYH3******Jeyl+yAvi" -> null
#      ~ sha256_fingerprint = "SHA256:LAU9OkL*****txJzNRas/*******60YoE" -> (known after apply)
#      + ssh_key_id         = "bc5277e1-0154-1234-80ef-1234556" -> (known after apply)
#    }

# Creating SSH key
resource "ruvds_ssh" "my_key2" {
  name       = "MacBook2"
  public_key = "ssh-ed25519 AAAAC**********nCC92urGnbwZ6iU/GuWYH3******Jeyl+yAvi"
}