# monitor for shipping APIs

> [!error] RIP
> this unfortunately no longer works
> due to shippingapimonitor not using
> that cloud flare captcha.

> [!example] be advised
> this is an example! it should contain useful information that the user should
> take into consideration before going forward!

#### Install:

`$ go install github.com/timmyha/shipmon`

#### Usage:

`$ shipmon`

#### Outcome:

```
+------------+--------+---------------+--------+------+-------+-------------------------+
|  CARRIER   | STATUS | RESPONSE TIME | <3MIN^ | <1H^ | <24H^ |       LAST UPDATE       |
+------------+--------+---------------+--------+------+-------+-------------------------+
| FedEx      | Online | 1.29 sec      | 100%   | 100% | 100%  | 2024-02-01 22:59:10 EST |
| UPS        | Online | 7.21 sec      | 100%   | 100% | 100%  | 2024-02-01 22:59:09 EST |
| USPS       | Online | 0.27 sec      | 100%   | 100% | 100%  | 2024-02-01 22:59:10 EST |
| CanadaPost | Online | 0.24 sec      | 100%   | 100% | 100%  | 2024-02-01 22:59:11 EST |
+------------+--------+---------------+--------+------+-------+-------------------------+
^ = uptime
```
