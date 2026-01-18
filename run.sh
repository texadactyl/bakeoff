LOGFILE=run.log
CSVFILE=run.csv
set -e
echo "lang,nloops,elapsed-secs,computed-pi" > $CSVFILE
unset JAVA_TOOL_OPTIONS
exec > >(tee "$LOGFILE") 2>&1
date
echo
echo =============== C ============================
gcc --version | head -n 1
gcc -O3 -march=native -mtune=native -ffast-math main.c
a.out >> $CSVFILE
echo
echo =============== D \(ldc2\) =====================
ldc2 --version | head -n 1
ldc2 -O3 -release -mcpu=native -ffast-math main.d -of=d.out
d.out >> $CSVFILE
echo
echo =============== zig ==========================
echo zig `zig version`
zig build-exe main.zig -O ReleaseFast -mcpu=native
mv main ziggy.out
ziggy.out >> $CSVFILE
echo
echo =============== Go ===========================
go version
go build -o go.out main.go
go.out >> $CSVFILE
echo
echo =============== Rust =========================
rustc --version
rustc -C opt-level=3 -o rs.out main.rs
rs.out >> $CSVFILE
echo
echo =============== Java Compile ================================================
javac --version
javac -Xlint -Werror main.java
echo =============== Hotspot JVM =============================================
java --version | head -n 1
java -server main >> $CSVFILE
echo
cat $CSVFILE
