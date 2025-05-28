# DocuSeal Discord Redirector

A webhook server that receives DocuSeal form events and forwards them to Discord as rich embeds.

## Features

- Receives DocuSeal webhook events (form viewed, started, completed, declined)
- Forwards events to Discord with formatted embeds
- Color-coded messages based on event type
- Includes signature images for completed forms
- API key authentication for security

## Example Discord Message

![Example Discord Message](docs/example.png)

## Setup

### 1. Configuration

Copy the example environment file and configure:

```bash
cp .env.example .env
```

Edit `.env` with your settings:

```env
PORT=8080
KEYS_FILE=keys.txt
DISCORD_WEBHOOK_URL=https://discord.com/api/webhooks/YOUR_WEBHOOK_URL
```

### 2. API Keys

Create or edit `keys.txt` with your authorized API keys (one per line):

```
your-secret-api-key-1
your-secret-api-key-2
```

Lines starting with `#` are treated as comments.

### 3. Build and Run

```bash
# Install dependencies
go mod download

# Build
go build -o main ./cmd/server/main.go

# Run
./main
```

Or use Air for development with hot reload:

```bash
air
```

## Usage

### Webhook Endpoint

The server exposes a webhook endpoint at:

```
POST /incoming/{key}
```

Where `{key}` must be one of the authorized keys from your `keys.txt` file.

### DocuSeal Configuration

Configure DocuSeal to send webhooks to:

```
https://your-domain.com/incoming/your-secret-api-key-1
```

### Supported Events

- **form.viewed** - Blue embed when someone views a form
- **form.started** - Yellow embed when someone starts filling a form
- **form.completed** - Green embed when a form is completed (includes signature)
- **form.declined** - Red embed when a form is declined

## License

This project is licensed under the GNU Affero General Public License v3.0.
