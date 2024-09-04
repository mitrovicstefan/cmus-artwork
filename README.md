# Cmus pixel artwork

![image](https://github.com/user-attachments/assets/78eb33d0-513e-43a8-bf89-fff469f22d1a)

Small [cmus](https://github.com/cmus/cmus) util for displaying artwork in pixelated form. Also allows pause/play/next/previous commands. The purpose is exclusively aesthetic.
Only tested and supported on linux and mac.

## Installation

Download the latest release using wget or curl

```
wget https://github.com/mitrovicstefan/cmus-artwork/releases/download/v1.0.0/cmus-artwork
```

```
curl https://github.com/mitrovicstefan/cmus-artwork/releases/download/v1.0.0/cmus-artwork
```

Move the file to your local binaries

```
sudo mv cmus-artwork /usr/local/bin/
```

Run it!

```
cmus-artwork
```

## Building and development

If you wish to build the binary yourself, it's extremely easy! First make sure that you have [go installed locally](https://go.dev/doc/install).

Once it's installed, clone the repository -

```
git clone https://github.com/mitrovicstefan/cmus-artwork
cd cmus-artwork
```

Run `main.go` to automatically install all the dependencies -

```
go run main.go
```

To build a binary run -

```
go build
```

From here you can start developing and messing around with the code. PRs are ALWAYS welcome!

## Feature requests and bug reports

To request a feature or report a bug open a new issue in this repo. I'll get to it in my free time :)

## Thank you
[moby/term](https://github.com/moby/term/tree/main) codebase for mac support\
[TheZoraiz/ascii-image-converter](https://github.com/TheZoraiz/ascii-image-converter) is library I use to convert images to ascii
