# bakeoff
Leibniz Algorithm for Pi on Multiple Language Platforms

### Recipe to Get Up and Running
* Install gcc.
* Install ldc2 version of D.
* Install zig.
* Install Go.
* Install Java version 21.
* Install Jacobin.

 ### Execution
 Bash script: run.sh

### Sample Log (run.log)
```
=============== C ============================
Number of rounds: 3000000000
Elapsed time (s): 3.637
3.14159265

=============== D (ldc2) =====================
Number of rounds: 3000000000
Elapsed time (s): 3.640
3.14159265

=============== zig ==========================
Number of rounds: 3000000000
Elapsed time (s): 3.638
3.14159265

=============== Go ===========================
Number of rounds: 3000000000
Elapsed time (s): 3.639
3.14159265

=============== Java Compile ================================================
=============== Hotspot -server =============================================
Number of rounds: 3000000000
Elapsed time (s): 8.005
3.14159265

=============== Hotspot -server -XX:+TieredCompilation ======================
Number of rounds: 3000000000
Elapsed time (s): 8.001
3.14159265

=============== Jacobin ======================
Number of rounds: 3000000000
Elapsed time (s): 7.996
3.14159265
```

### Afterthoughts

Feel free to hack run.sh.
Issues: Suggestions, Improvements, and questions are most welcome.
