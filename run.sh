#!/usr/bin/env bash

set -euo pipefail
unset JAVA_TOOL_OPTIONS
export rounds=3000000000 # 3 billion

LOGFILE=run.log
exec > >(tee "$LOGFILE") 2>&1
CSVFILE=run.csv
echo "lang,nloops,elapsed-secs,computed-pi" > $CSVFILE

timestamp() {
    date -u +"%Y-%m-%dT%H:%M:%SZ"
}

section() {
    echo
    echo "=== $1 ==="
}

kv() {
    printf "%-24s %s\n" "$1:" "$2"
}

timestamp
section "ENVIRONMENT"

kv "OS"        "${RUNNER_OS:-unknown}"
kv "ARCH"      "$(uname -m)"
kv "KERNEL"    "$(uname -sr)"

ARCH=$(uname -m)

# ----------------------------
# Architecture-specific flags
# ----------------------------
# Use GitHub Actions RUNNER_ARCH if available, fallback to uname
ARCH="${RUNNER_ARCH:-$(uname -m)}"

# Normalize to lowercase for consistent matching
ARCH=$(echo "$ARCH" | tr '[:upper:]' '[:lower:]')

case "$ARCH" in
    x64|x86_64|amd64)
        C_CPU="-march=native -mtune=native"
        D_CPU="-mcpu=native"
        ZIG_CPU="-mcpu=native"
        RUST_CPU="native"
        ;;
    arm64|aarch64)
        C_CPU="-march=native"
        D_CPU="-mcpu=native"
        ZIG_CPU="-mcpu=native"
        RUST_CPU="native"
        ;;
    *)
        echo "WARNING: Unknown architecture '$ARCH', using baseline"
        C_CPU="-march=native"
        D_CPU="-mcpu=native"
        ZIG_CPU="-mcpu=baseline"
        RUST_CPU="native"
        ;;
esac

COMMON_CFLAGS="-O3 -ffast-math"
COMMON_DFLAGS="-O3 -release -ffast-math"
COMMON_ZIGFLAGS="-O ReleaseFast"
COMMON_RUSTFLAGS="-C opt-level=3"

# ============================
# C
# ============================
section "C"

kv "Compiler" "$(gcc -dumpmachine)"
kv "Version"  "$(gcc --version | head -n 1)"

gcc $COMMON_CFLAGS $C_CPU main.c  -lm -o c.out
./c.out | tee -a $CSVFILE

# ============================
# D (LDC2)
# ============================
section "D (ldc2)"

kv "Compiler" "ldc2"
kv "Version"  "$(ldc2 --version 2>&1 | grep -m1 'LDC' | tr '\n' ' ')"

ldc2 $COMMON_DFLAGS $D_CPU main.d -of=d.out
./d.out | tee -a $CSVFILE

# ============================
# Go
# ============================
section "Go"

kv "Compiler" "go"
kv "Version"  "$(go version)"

go build -trimpath -ldflags="-s -w" -gcflags="-B" -o go.out main.go
./go.out | tee -a $CSVFILE

# ============================
# Java
# ============================
section "Java Compile"

kv "javac" "$(javac --version)"
javac -Xlint -Werror main.java

section "HotSpot JVM"
kv "java" "$(java --version | head -n 1)"
java -server main | tee -a $CSVFILE


# ============================
# Nim
# ============================
section "Nim"

kv "Compiler" "nim"
kv "Version"  "$(nim --version | head -n 1)"

nim c -d:release -d:danger --opt:speed --passC:-march=native -o:nim.out --hints:off main.nim
nim.out  | tee -a $CSVFILE


# ============================
# Rust
# ============================
section "Rust"

kv "Compiler" "rustc"
kv "Version"  "$(rustc --version)"

rustc -C opt-level=3 -C target-cpu=native -o rs.out main.rs
./rs.out | tee -a $CSVFILE

# ============================
# Zig
# ============================
section "Zig"

kv "Compiler" "zig"
kv "Version"  "$(zig version)"

zig build-exe main.zig \
    $COMMON_ZIGFLAGS \
    $ZIG_CPU \
    -femit-bin=zig.out

./zig.out | tee -a $CSVFILE

# ============================
# Show CSV file
# ============================
section "Show CSV file"
echo
cat $CSVFILE

section "DONE"
timestamp


