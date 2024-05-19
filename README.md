
# Menpo

_Proof of Concept for file encryption/decryption/hash._



- Generate sha256sum of local file.
- Encrypt local file.
  - Symmetric authenticated encryption.
  - AES GCM.
  - Cryptographically secure pseudo random 12 byte nonce.
  - Cryptographically secure pseudo random 16 byte secret key.

Maybe will add asymmetric encryption option in future.
Maybe illustrated my point already and it's not needed.

![selecting file](/screenshots/select_file.png "selecting file")
![shasum](/screenshots/shasum.png "shasum")
![encrypting](/screenshots/encrypting.png "encrypting")
![decrypting](/screenshots/decrypting.png "decrypting")
