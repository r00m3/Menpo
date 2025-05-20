# Menpo
###### _Proof of Concept for file encryption/decryption/hash._

<br/>

- Generate sha256sum of local file.
- Encrypt local file.
- Decrypt local file.
	- Symmetric authenticated encryption.
	- AES GCM.
	- Cryptographically secure pseudo random 12 byte nonce.
	- Cryptographically secure pseudo random 16 byte secret key.

![selecting file](/screenshots/select_file.png "selecting file")

<br/>

![shasum](/screenshots/shasum.png "shasum")

<br/>

![encrypting](/screenshots/encrypting.png "encrypting")

<br/>

![decrypting](/screenshots/decrypting.png "decrypting")
###### _Maybe will add asymmetric encryption option in future. Maybe illustrated my point already and it's not needed._
