# certgen
_Generate and install local certificate_

Easy. Handy. Free.

## The Idea
This fork strips away the server part of the original repo. So essentially, this repo will only create and install the certificates
but will not start a http server. Refer to the original repo called [sserve](https://github.com/daquinoaldo/sserve) if this isn't what you're looking for.

### Warning
The `rootCA-key.pem` file that mkcert automatically generates when installing sserve gives complete power to intercept secure requests from your machine. Do not share it.

### License
Is released under [AGPL-3.0 - GNU Affero General Public License v3.0](LICENSE).

#### Briefly:
- modification and redistribution allowed for both private and **commercial use**
- you must **grant patent right to the owner and to all the contributors**
- you must **keep it open source** and distribute under the **same license**
- changes must be documented
- include a limitation of liability and it **does not provide any warranty**

### Warranty
THIS TOOL IS PROVIDED "AS IS" WITHOUT WARRANTY OF ANY KIND.
THE ENTIRE RISK AS TO THE QUALITY AND PERFORMANCE OF THE PROGRAM IS WITH YOU.
For the full warranty check the [LICENSE](LICENSE).