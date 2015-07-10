## State types

* **SOFT** - at least one check changed state but under retries/attempts count, or just after state change
* **HARD** - amount of failed checks >= `max_check_attempts`

## Service state

* 0 - **OK**
* 1 - **WARNING**
* 2 - **CRITICAL**
* 3+ - **UNKNOWN** - check failed for reason other than "service/host is not working"; example would be missing deps to run check or passive check timeout

## Host state

0 - **UP**
1 - **DOWN**
