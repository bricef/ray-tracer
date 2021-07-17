#!/bin/bash
 
FRAME_DIRECTORY="$1" 
OUTPUT_GIF_NAME="$2"
ffmpeg -f image2 -i "${FRAME_DIRECTORY}/frame-%d.png" tmp.avi
ffmpeg -i tmp.avi -pix_fmt rgb24 "${OUTPUT_GIF_NAME}.gif"