# Meshtastic Compression Showdown

This project contain benchmarks of various compression algorithms applied on a data of meshtastic packets.

For context a Reciprocal Compression Ratio **above** 1 means the compressed data is **bigger** than the uncompressed data.
One **bellow** 1 means the compressed data is **smaller** than the uncompressed data.

## Results

| Compressor | Average Reciprocal Compression Ratio | Average Reciprocal Compression Ratio (TEXT_MESSAGE_APP only) |
|------------|--------------------------------------|-------------------------------------------------------------|
| `unishox2_meshtastic` | 0.9983 | 0.7637 |
| `noop` | 1.0000 | 1.0000 |
| `flate_klauspost` | 1.0647 | 1.0595 |
| `flate_std` | 1.1275 | 1.1595 |
| `lzw_std` | 1.1478 | 1.1330 |
| `zstd_klauspost` | 1.1792 | 1.2531 |
| `zlib_klauspost` | 1.1868 | 1.2543 |
| `zlib_std` | 1.2497 | 1.3544 |
| `s2_klauspost` | 1.3629 | 1.5730 |
| `snappy_klauspost` | 1.3646 | 1.5775 |
| `gzip_klauspost` | 1.4311 | 1.6440 |
| `gzip_std` | 1.4940 | 1.7440 |
## CDF Graphs

The following graphs show the cumulative distribution function (CDF) of the reciprocal compression ratios for each compressor.

### `unishox2_meshtastic`

![unishox2_meshtastic CDF](graphs/unishox2_meshtastic_cdf.png)

![unishox2_meshtastic only TEXT_MESSAGE_APP CDF](graphs/unishox2_meshtastic only TEXT_MESSAGE_APP_cdf.png)

### `noop`

![noop CDF](graphs/noop_cdf.png)

![noop only TEXT_MESSAGE_APP CDF](graphs/noop only TEXT_MESSAGE_APP_cdf.png)

### `flate_klauspost`

![flate_klauspost CDF](graphs/flate_klauspost_cdf.png)

![flate_klauspost only TEXT_MESSAGE_APP CDF](graphs/flate_klauspost only TEXT_MESSAGE_APP_cdf.png)

### `flate_std`

![flate_std CDF](graphs/flate_std_cdf.png)

![flate_std only TEXT_MESSAGE_APP CDF](graphs/flate_std only TEXT_MESSAGE_APP_cdf.png)

### `lzw_std`

![lzw_std CDF](graphs/lzw_std_cdf.png)

![lzw_std only TEXT_MESSAGE_APP CDF](graphs/lzw_std only TEXT_MESSAGE_APP_cdf.png)

### `zstd_klauspost`

![zstd_klauspost CDF](graphs/zstd_klauspost_cdf.png)

![zstd_klauspost only TEXT_MESSAGE_APP CDF](graphs/zstd_klauspost only TEXT_MESSAGE_APP_cdf.png)

### `zlib_klauspost`

![zlib_klauspost CDF](graphs/zlib_klauspost_cdf.png)

![zlib_klauspost only TEXT_MESSAGE_APP CDF](graphs/zlib_klauspost only TEXT_MESSAGE_APP_cdf.png)

### `zlib_std`

![zlib_std CDF](graphs/zlib_std_cdf.png)

![zlib_std only TEXT_MESSAGE_APP CDF](graphs/zlib_std only TEXT_MESSAGE_APP_cdf.png)

### `s2_klauspost`

![s2_klauspost CDF](graphs/s2_klauspost_cdf.png)

![s2_klauspost only TEXT_MESSAGE_APP CDF](graphs/s2_klauspost only TEXT_MESSAGE_APP_cdf.png)

### `snappy_klauspost`

![snappy_klauspost CDF](graphs/snappy_klauspost_cdf.png)

![snappy_klauspost only TEXT_MESSAGE_APP CDF](graphs/snappy_klauspost only TEXT_MESSAGE_APP_cdf.png)

### `gzip_klauspost`

![gzip_klauspost CDF](graphs/gzip_klauspost_cdf.png)

![gzip_klauspost only TEXT_MESSAGE_APP CDF](graphs/gzip_klauspost only TEXT_MESSAGE_APP_cdf.png)

### `gzip_std`

![gzip_std CDF](graphs/gzip_std_cdf.png)

![gzip_std only TEXT_MESSAGE_APP CDF](graphs/gzip_std only TEXT_MESSAGE_APP_cdf.png)

