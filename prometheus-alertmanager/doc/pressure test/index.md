
https://stackoverflow.com/questions/2925606/how-to-create-a-cpu-spike-with-a-bash-command


## example 1

You can also do

`dd if=/dev/zero of=/dev/null`
To run more of those to put load on more cores, try to fork it:

`fulload() { dd if=/dev/zero of=/dev/null | dd if=/dev/zero of=/dev/null | dd if=/dev/zero of=/dev/null | dd if=/dev/zero of=/dev/null & }; fulload; read; killall dd`


Repeat the command in the curly brackets as many times as the number of threads you want to produce (here 4 threads). Simple enter hit will stop it (just make sure no other dd is running on this user or you kill it too).


## example 2

Example to stress 2 cores for 60 seconds

`stress --cpu 2 --timeout 60`

## example 3

`stress -c 40`
to stress all your CPUs (however you have) with 40 threads each running a complex sqrt computation on a ramdomly generated numbers.

You can even define the timeout of the program

`stress -c 40 -timeout 10s`