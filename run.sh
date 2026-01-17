LOGFILE=run.log
set -e
unset JAVA_TOOL_OPTIONS

exec > >(tee "$LOGFILE") 2>&1
echo =============== C ============================
gcc -O3 -march=native -mtune=native -ffast-math main.c
a.out
echo
echo =============== D \(ldc2\) =====================
ldc2 -O3 -release -mcpu=native -ffast-math main.d -of=d.out
d.out
echo
echo =============== zig ==========================
zig build-exe main.zig -O ReleaseFast -mcpu=native
mv main ziggy.out
ziggy.out
echo
echo =============== Go ===========================
go build -o go.out main.go
go.out
echo
echo =============== Java Compile ================================================
javac -Xlint -Werror main.java
echo =============== Hotspot -server =============================================
java -server main
echo
echo =============== Hotspot -server -XX:+TieredCompilation ======================
java -server -XX:+TieredCompilation main
echo
echo =============== Jacobin ======================
java main
