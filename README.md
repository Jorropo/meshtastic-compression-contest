# Meshtastic Compression Showdown

This project contain benchmarks of various compression algorithms applied on a data of meshtastic packets.

For context a Reciprocal Compression Ratio **above** 1 means the compressed data is **bigger** than the uncompressed data.
One **bellow** 1 means the compressed data is **smaller** than the uncompressed data.

## Results

| Compressor | Average Reciprocal Compression Ratio |
|------------|--------------------------------------|
| `unishox2_meshtastic` | 0.9983 |
| `noop` | 1.0000 |
| `flate_klauspost` | 1.0647 |
| `flate_std` | 1.1275 |
| `lzw_std` | 1.1478 |
| `zstd_klauspost` | 1.1792 |
| `zlib_klauspost` | 1.1868 |
| `zlib_std` | 1.2497 |
| `s2_klauspost` | 1.3629 |
| `snappy_klauspost` | 1.3646 |
| `gzip_klauspost` | 1.4312 |
| `gzip_std` | 1.4940 |
## CDF Graphs

The following graphs show the cumulative distribution function (CDF) of the reciprocal compression ratios for each compressor.

### `unishox2_meshtastic` 0.9983

![unishox2_meshtastic CDF](graphs/unishox2_meshtastic_cdf.png)

### `noop` 1.0000

![noop CDF](graphs/noop_cdf.png)

### `flate_klauspost` 1.0647

![flate_klauspost CDF](graphs/flate_klauspost_cdf.png)

### `flate_std` 1.1275

![flate_std CDF](graphs/flate_std_cdf.png)

### `lzw_std` 1.1478

![lzw_std CDF](graphs/lzw_std_cdf.png)

### `zstd_klauspost` 1.1792

![zstd_klauspost CDF](graphs/zstd_klauspost_cdf.png)

### `zlib_klauspost` 1.1868

![zlib_klauspost CDF](graphs/zlib_klauspost_cdf.png)

### `zlib_std` 1.2497

![zlib_std CDF](graphs/zlib_std_cdf.png)

### `s2_klauspost` 1.3629

![s2_klauspost CDF](graphs/s2_klauspost_cdf.png)

### `snappy_klauspost` 1.3646

![snappy_klauspost CDF](graphs/snappy_klauspost_cdf.png)

### `gzip_klauspost` 1.4312

![gzip_klauspost CDF](graphs/gzip_klauspost_cdf.png)

### `gzip_std` 1.4940

![gzip_std CDF](graphs/gzip_std_cdf.png)

