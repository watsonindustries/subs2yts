# About

`subs2yts` is a small tool to automatically convert raw subtitle files generated from YouTube VODs/Livestreams into YouTube comment timestamp formatted lines.

## Example

```bash
subs2yts -i resampled-2UtQM5LMopI.wav.vtt -o out.txt
```

Will produce something like:

```
...
00:03:55 I actually think it would be okay if I had the super chat sign, but I don't want to risk getting in trouble.
00:04:03 Okay, so the game's a little lower quality. I'm actually streaming at 1080p and then the game is at was running at like 720p resolution. So what do you guys think?
00:04:15 How does it look? It's not it's not beautiful case. I think some fog has been turned off. It looks like
...
```

Then you can simply copy and paste the lines into the comment section of the target YouTube video.

## Build

Install Go toolchain and run `go build .`

## Install

You can install the binary to your `$GOPATH/bin` by running `go install .`

## Notes
