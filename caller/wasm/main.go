package main

import (
	"context"
	"fmt"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"os"
)

func main() {
	wasmBytes, err := os.ReadFile("wasm.wasm")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	mod, err := r.Instantiate(ctx, wasmBytes)
	if err != nil {
		panic(err)
	}

	fn := mod.ExportedFunction("convert")
	malloc := mod.ExportedFunction("malloc")
	free := mod.ExportedFunction("free")

	param := "wasm"
	results, err := malloc.Call(ctx, uint64(len(param)))
	if err != nil {
		panic(err)
	}
	paramPtr, paramSize := results[0], uint64(len(param))
	defer free.Call(ctx, paramPtr)

	if !mod.Memory().Write(uint32(paramPtr), []byte(param)) {
		panic(fmt.Errorf("write memory pointer %d failed", paramPtr))
	}

	res, err := fn.Call(ctx, paramPtr, paramSize)
	if err != nil {
		panic(err)
	}

	resPtr := uint32(res[0] >> 32)
	resSize := uint32(res[0])
	if resString, ok := mod.Memory().Read(resPtr, resSize); !ok {
		panic(fmt.Errorf("read memory pointer %d failed", resPtr))
	} else {
		fmt.Println(string(resString))
	}

}
