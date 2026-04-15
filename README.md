<table>
  <tr>
    <td><img src="public/img/galaxy.png" width="50" alt="Sonicary"></td>
    <td><h1>Sonicary</h1></td>
  </tr>
</table>

Sonicary is a self-hosted music library manager focused on importing, organizing, and maintaining music collections.

It began as a fork of [Soulsolid](https://github.com/contre95/soulsolid), building on the excellent foundation created by its original author, [Contre](https://contre.io/). Sonicary is taking the project in a different direction, with a stronger focus on safer library setup, organization, and long-term collection management for self-hosted use.

## Attribution

Sonicary is based on [Soulsolid](https://github.com/contre95/soulsolid). Full credit and appreciation go to [Contre](https://contre.io/) for creating the original project and laying the groundwork this fork builds on.

This fork does not attempt to erase or obscure that origin. It exists to explore a different workflow and product direction.

## Screenshots

<table>
  <tr>
    <td>
      <img src="./docs/screen0.jpg" />
    </td>
    <td>
      <img src="./docs/screen1.jpg" />
    </td>
  </tr>
</table>

## Features

- **Music Library Management**: Organize and browse albums, artists, and tracks
- **Downloading**: Download tracks and albums
- **Importing**: Import new music from intake directories with automatic fingerprinting
- **Metadata Tagging**: Auto-tag using MusicBrainz and Discogs APIs
- **Telegram Integration**: Control via Telegram bot
- **Web UI**: Mobile-friendly interface for library and import workflows
- **Job Management**: Background processing for downloads, imports, lyrics, and related tasks

> Documentation and demo links are being reworked for Sonicary and are not yet published.

## Quick Start

### 🦭 Container Usage

The application can run without copying `config.yaml` into the container. If no config file exists, it will automatically create one with sensible defaults.

#### Environment Variable Support

Sonicary supports environment variables in configuration files using the `!env_var` tag:

```yaml
telegram:
  token: !env_var TELEGRAM_BOT_TOKEN
metadata:
  providers:
    discogs:
      secret: !env_var DISCOGS_API_KEY
```

The application will fail to start if a referenced environment variable is not set.

```bash
# Build the image
podman build -t sonicary .

# Run with environment variables
podman run -d   --name sonicary   -p 3535:3535   -v /host/music:/app/library   -v /host/downloads:/app/downloads   -v /host/logs:/app/logs   -v /host/library.db:/data/library.db   -v /host/config.yaml:/config/config.yaml   sonicary
```

- `/app/library` is the managed music library location
- `/app/downloads` is the intake/download path used for new imports

The web interface will be available at `http://localhost:3535`.

## Development

To set up the development environment:

### Option 1: Manual Setup

```bash
cp config.example.yaml config.yaml
npm run dev
go run ./src/main.go
```

### Option 2: Using Nix

If you have Nix installed, use the provided `dev.nix` shell:

```bash
nix-shell dev.nix
go run ./src/main.go
```

The web interface will be available at `http://localhost:3535`.
