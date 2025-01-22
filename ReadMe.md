# Go Blockchain Implementation

```markdown

A minimal yet functional blockchain implementation in Go featuring Proof-of-Work consensus, persistent storage, and REST API endpoints.

## Features

- 📦 Block structure with cryptographic linking
- ⛏️ Configurable Proof-of-Work mining
- 🔗 Immutable blockchain ledger
- 💾 JSON-based persistence
- 🌐 REST API for chain interaction
- ✅ Comprehensive test suite

```

## Configuration

Edit `config.json` to customize blockchain parameters:

```json
{
  "difficulty": 4,
  "block_reward": 10,
  "genesis_data": "Genesis Block",
  "port": ":8080"
}

```

| Parameter | Description | Default |
| --- | --- | --- |
| `difficulty` | Leading zeros required in block hash | 4 |
| `block_reward` | Miner reward (future use) | 10 |
| `genesis_data` | Initial block data | "Genesis Block" |
| `port` | API server port | ":8080" |

## Project Structure

```
blockchain/
├── blockchain/         # Core implementation
│   ├── block.go        # Block structure and hashing
│   ├── blockchain.go   # Chain management
│   ├── config.go       # Configuration loader
│   └── pow.go          # Proof-of-Work system
├── tests/              # Test suite
│   └── blockchain_test.go
├── config.json         # Runtime parameters
└── main.go             # API server entrypoint

```

## API Endpoints

### GET `/chain`

Returns complete blockchain

**Response**:

```json
[
  {
    "index": 0,
    "timestamp": "2023-07-20T12:34:56Z",
    "transactions": [...],
    "prev_hash": "0",
    "hash": "0000a8c2...",
    "nonce": 12345,
    "miner_address": "system"
  }
]

```

### GET `/mine`

Creates and mines new block

**Response**:

```json
{
  "index": 1,
  "timestamp": "2023-07-20T12:35:07Z",
  "transactions": [],
  "prev_hash": "0000a8c2...",
  "hash": "0000e3b0...",
  "nonce": 45678,
  "miner_address": "test_miner"
}

```

## Mining Process

1. New block created with:
    - Previous block's hash
    - Current timestamp
    - Empty transaction list
    - Incremented index
2. Proof-of-Work execution:

    ```go
    func (pow *ProofOfWork) Mine() (int, string) {
        target := strings.Repeat("0", pow.Config.Difficulty)
        for nonce := 0; ; nonce++ {
            pow.Block.Nonce = nonce
            hash := pow.Block.CalculateHash()
            if strings.HasPrefix(hash, target) {
                return nonce, hash
            }
        }
    }

    ```

3. Validated block added to chain
4. Chain state persisted to `blockchain.json`

## Persistence

Blocks are saved in human-readable JSON format:

```bash
cat blockchain.json

```

Example output:

```json
[

   {

      "index": 0,

      "timestamp": "2025-01-22T23:31:45+03:00",

      "transactions": [

         {

            "Sender": "system",

            "Recipient": "genesis",

            "Amount": 0

         }

      ],

      "prev_hash": "0",

      "hash": "006a321e846d2d78b795c4fa7b8d823bad9509e5e4c8c825a0dc32d3a3b9e241",

      "nonce": 0,

      "miner_address": "system"

   },

   {

      "index": 1,

      "timestamp": "2025-01-22T23:32:46+03:00",

      "transactions": null,

      "prev_hash": "006a321e846d2d78b795c4fa7b8d823bad9509e5e4c8c825a0dc32d3a3b9e241",

      "hash": "0000d7a205c097aa024f2e4ace2657b79739d40b10b9a20efa44c102cb36ef24",

      "nonce": 213146,

      "miner_address": "test_miner"

   }

]

```

## Testing

Run validation suite:

```bash
go test -v ./tests/...

```

Test coverage includes:

- Genesis block initialization
- Block addition workflow
- POW validation mechanics
- Chain integrity checks

Example test output:

```
=== RUN   TestNewBlockchain
--- PASS: TestNewBlockchain (0.00s)
=== RUN   TestAddBlock
--- PASS: TestAddBlock (0.32s)
=== RUN   TestPOWValidation
--- PASS: TestPOWValidation (0.18s)
PASS

```

## Security Considerations

1. **Cryptographic Integrity**
    - SHA-256 hashing
    - Immutable block links
    - POW difficulty enforcement
2. **Network Protection**
    - Basic HTTP server (add HTTPS in production)
    - Input validation recommended for production use
3. **Data Protection**
    - File system persistence
    - Atomic write operations

## Future Extensions

1. Transaction validation system
2. Peer-to-peer networking
3. Wallet/address system
4. UTXO model implementation
5. Consensus algorithm upgrades
6. Smart contract support

---

**Version**: 1.0.0-alpha
