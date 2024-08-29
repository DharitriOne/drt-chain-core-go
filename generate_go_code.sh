source ~/.profile

# data/batch
protoc \
    -I="$HOME/go/src" \
    -I="$HOME/go/src/github.com/DharitriOne/protobuf" \
    -I="$PWD/data/batch" \
    --gogoslick_out="$PWD/data/batch" \
    $PWD/data/batch/*.proto

# data/dcdt
protoc \
    -I="$HOME/go/src" \
    -I="$HOME/go/src/github.com/DharitriOne/protobuf" \
    -I="$PWD/data/dcdt" \
    --gogoslick_out="$PWD/data/dcdt" \
    $PWD/data/dcdt/proto/dcdt.proto

mv $PWD/data/dcdt/proto/dcdt.pb.go $PWD/data/dcdt/

# data/guardians
protoc \
    -I="$HOME/go/src" \
    -I="$HOME/go/src/github.com/DharitriOne/protobuf" \
    -I="$PWD/data/guardians" \
    --gogoslick_out="$PWD/data/guardians" \
    $PWD/data/guardians/*.proto

# data/block
protoc \
    -I="$HOME/go/src" \
    -I="$HOME/go/src/github.com/DharitriOne/protobuf" \
    -I="$PWD/data/block" \
    --gogoslick_out="$PWD/data/block" \
    $PWD/data/block/*.proto

# data/metrics
protoc \
    -I="$HOME/go/src" \
    -I="$HOME/go/src/github.com/DharitriOne/protobuf" \
    -I="$PWD/data/metrics" \
    --gogoslick_out="$PWD/data/metrics" \
    $PWD/data/metrics/*.proto

# data/validator
protoc \
    -I="$HOME/go/src" \
    -I="$HOME/go/src/github.com/DharitriOne/protobuf" \
    -I="$PWD/data/validator" \
    --gogoslick_out="$PWD/data/validator" \
    $PWD/data/validator/*.proto

# marshal/testSizeCheckUnmarshal
protoc \
    -I="$HOME/go/src" \
    -I="$HOME/go/src/github.com/DharitriOne/protobuf" \
    -I="$PWD/marshal/testSizeCheckUnmarshal" \
    --gogoslick_out="$PWD/marshal/testSizeCheckUnmarshal" \
    $PWD/marshal/testSizeCheckUnmarshal/*.proto

# data/alteredAccount
protoc \
    -I="$HOME/go/src" \
    -I="$HOME/go/src/github.com/DharitriOne/protobuf" \
    -I="$PWD/data/alteredAccount" \
    --gogoslick_out="$PWD/data/alteredAccount" \
    $PWD/data/alteredAccount/*.proto

mv $PWD/data/alteredAccount/github.com/DharitriOne/drt-chain-core-go/data/alteredAccount/alteredAccount.pb.go $PWD/data/alteredAccount/
rm -r $PWD/data/alteredAccount/github.com

# data/receipt
protoc \
    -I="$HOME/go/src" \
    -I="$HOME/go/src/github.com/DharitriOne/protobuf" \
    -I="$PWD/data/receipt" \
    --gogoslick_out="$PWD/data/receipt" \
    $PWD/data/receipt/*.proto

mv $PWD/data/receipt/github.com/DharitriOne/drt-chain-core-go/data/receipt/receipt.pb.go $PWD/data/receipt/
rm -r $PWD/data/receipt/github.com

# data/rewardTx
protoc \
    -I="$HOME/go/src" \
    -I="$HOME/go/src/github.com/DharitriOne/protobuf" \
    -I="$PWD/data/rewardTx" \
    --gogoslick_out="$PWD/data/rewardTx" \
    $PWD/data/rewardTx/*.proto

mv $PWD/data/rewardTx/github.com/DharitriOne/drt-chain-core-go/data/rewardTx/rewardTx.pb.go $PWD/data/rewardTx/
rm -r $PWD/data/rewardTx/github.com

# data/smartContractResult
protoc \
    -I="$HOME/go/src" \
    -I="$HOME/go/src/github.com/DharitriOne/protobuf" \
    -I="$PWD/data/smartContractResult" \
    --gogoslick_out="$PWD/data/smartContractResult" \
    $PWD/data/smartContractResult/*.proto

mv $PWD/data/smartContractResult/github.com/DharitriOne/drt-chain-core-go/data/smartContractResult/smartContractResult.pb.go $PWD/data/smartContractResult/
rm -r $PWD/data/smartContractResult/github.com

# data/transaction
protoc \
    -I="$HOME/go/src" \
    -I="$HOME/go/src/github.com/DharitriOne/protobuf" \
    -I="$PWD/data/transaction" \
    --gogoslick_out="$PWD/data/transaction" \
    $PWD/data/transaction/*.proto

mv $PWD/data/transaction/github.com/DharitriOne/drt-chain-core-go/data/transaction/*.pb.go $PWD/data/transaction/
rm -r $PWD/data/transaction/github.com

# data/outport
protoc \
    -I="$HOME/go/src" \
    -I="$HOME/go/src/github.com/DharitriOne/protobuf" \
    -I="$PWD/data/outport" \
    --gogoslick_out="$PWD/data/outport" \
    $PWD/data/outport/*.proto

mv $PWD/data/outport/github.com/DharitriOne/drt-chain-core-go/data/outport/*.pb.go $PWD/data/outport/
rm -r $PWD/data/outport/github.com