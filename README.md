# Meshtastic Compression Showdown

This project contain benchmarks of various compression algorithms applied on a data of meshtastic packets.

For context a Reciprocal Compression Ratio **above** 1 means the compressed data is **bigger** than the uncompressed data.
One **bellow** 1 means the compressed data is **smaller** than the uncompressed data.

## Results

| Compressor | Average Reciprocal Compression Ratio (TEXT_MESSAGE_APP only) |
|------------|--------------------------------------------------------------|
| `shoco_TextEn_tmthrgd_Jorropo` | 0.7711 |
| `unishox2_meshtastic` | 0.7892 |
| `smaz_cespare_Jorropo` | 0.8253 |
| `snowflake_Jorropo` | 0.8253 |
| `shoco_TextEn_tmthrgd` | 0.8434 |
| `shoco_WordsEn_tmthrgd_Jorropo` | 0.8614 |
| `shoco_FilePath_tmthrgd_Jorropo` | 0.8675 |
| `smaz_cespare` | 0.9036 |
| `shoco_Emails_tmthrgd_Jorropo` | 0.9096 |
| `shoco_WordsEn_tmthrgd` | 0.9337 |
| `shoco_FilePath_tmthrgd` | 0.9398 |
| `shoco_Emails_tmthrgd` | 0.9819 |
| `noop` | 1.0000 |
| `zstd_klauspost_chopped_Jorropo` | 1.0361 |
| `lz4_cloudflare` | 1.0602 |
| `lz4_cloudflareHC` | 1.0602 |
| `flate_klauspost` | 1.1145 |
| `rle_inkyblackness` | 1.1325 |
| `lzw_std` | 1.1627 |
| `flate_std` | 1.2289 |
| `zstd_klauspost` | 1.3253 |
| `zlib_klauspost` | 1.3313 |
| `zlib_std` | 1.4458 |
| `lz4_pierrec` | 1.5422 |
| `s2_klauspost` | 1.6506 |
| `snappy_klauspost` | 1.6506 |
| `gzip_klauspost` | 1.7651 |
| `gzip_std` | 1.8795 |

| Compressor | Average Reciprocal Compression Ratio |
|------------|--------------------------------------|
| `shoco_TextEn_tmthrgd_Jorropo` | 0.9991 |
| `unishox2_meshtastic` | 0.9992 |
| `smaz_cespare_Jorropo` | 0.9993 |
| `snowflake_Jorropo` | 0.9993 |
| `shoco_WordsEn_tmthrgd_Jorropo` | 0.9995 |
| `shoco_FilePath_tmthrgd_Jorropo` | 0.9995 |
| `shoco_Emails_tmthrgd_Jorropo` | 0.9997 |
| `noop` | 1.0000 |
| `zstd_klauspost_chopped_Jorropo` | 1.0181 |
| `lz4_cloudflareHC` | 1.0330 |
| `lz4_cloudflare` | 1.0335 |
| `flate_klauspost` | 1.0650 |
| `rle_inkyblackness` | 1.0843 |
| `flate_std` | 1.1276 |
| `lzw_std` | 1.1452 |
| `shoco_TextEn_tmthrgd` | 1.1714 |
| `shoco_WordsEn_tmthrgd` | 1.1723 |
| `shoco_Emails_tmthrgd` | 1.1736 |
| `shoco_FilePath_tmthrgd` | 1.1749 |
| `zstd_klauspost` | 1.1835 |
| `zlib_klauspost` | 1.1901 |
| `zlib_std` | 1.2527 |
| `smaz_cespare` | 1.2528 |
| `lz4_pierrec` | 1.3095 |
| `s2_klauspost` | 1.3710 |
| `snappy_klauspost` | 1.3743 |
| `gzip_klauspost` | 1.4404 |
| `gzip_std` | 1.5030 |

## CDF Graphs

The following graphs show the cumulative distribution function (CDF) of the reciprocal compression ratios for each compressor.

### `shoco_TextEn_tmthrgd_Jorropo`

![shoco_TextEn_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_TextEn_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_TextEn_tmthrgd_Jorropo CDF](graphs/shoco_TextEn_tmthrgd_Jorropo_cdf.png)

### `unishox2_meshtastic`

![unishox2_meshtastic only TEXT_MESSAGE_APP CDF](graphs/unishox2_meshtastic_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_meshtastic CDF](graphs/unishox2_meshtastic_cdf.png)

### `smaz_cespare_Jorropo`

![smaz_cespare_Jorropo only TEXT_MESSAGE_APP CDF](graphs/smaz_cespare_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![smaz_cespare_Jorropo CDF](graphs/smaz_cespare_Jorropo_cdf.png)

### `snowflake_Jorropo`

![snowflake_Jorropo only TEXT_MESSAGE_APP CDF](graphs/snowflake_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![snowflake_Jorropo CDF](graphs/snowflake_Jorropo_cdf.png)

### `shoco_WordsEn_tmthrgd_Jorropo`

![shoco_WordsEn_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_WordsEn_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_WordsEn_tmthrgd_Jorropo CDF](graphs/shoco_WordsEn_tmthrgd_Jorropo_cdf.png)

### `shoco_FilePath_tmthrgd_Jorropo`

![shoco_FilePath_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_FilePath_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_FilePath_tmthrgd_Jorropo CDF](graphs/shoco_FilePath_tmthrgd_Jorropo_cdf.png)

### `shoco_Emails_tmthrgd_Jorropo`

![shoco_Emails_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_Emails_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_Emails_tmthrgd_Jorropo CDF](graphs/shoco_Emails_tmthrgd_Jorropo_cdf.png)

### `noop`

![noop only TEXT_MESSAGE_APP CDF](graphs/noop_only_TEXT_MESSAGE_APP_cdf.png)

![noop CDF](graphs/noop_cdf.png)

### `zstd_klauspost_chopped_Jorropo`

![zstd_klauspost_chopped_Jorropo only TEXT_MESSAGE_APP CDF](graphs/zstd_klauspost_chopped_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![zstd_klauspost_chopped_Jorropo CDF](graphs/zstd_klauspost_chopped_Jorropo_cdf.png)

### `lz4_cloudflareHC`

![lz4_cloudflareHC only TEXT_MESSAGE_APP CDF](graphs/lz4_cloudflareHC_only_TEXT_MESSAGE_APP_cdf.png)

![lz4_cloudflareHC CDF](graphs/lz4_cloudflareHC_cdf.png)

### `lz4_cloudflare`

![lz4_cloudflare only TEXT_MESSAGE_APP CDF](graphs/lz4_cloudflare_only_TEXT_MESSAGE_APP_cdf.png)

![lz4_cloudflare CDF](graphs/lz4_cloudflare_cdf.png)

### `flate_klauspost`

![flate_klauspost only TEXT_MESSAGE_APP CDF](graphs/flate_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

![flate_klauspost CDF](graphs/flate_klauspost_cdf.png)

### `rle_inkyblackness`

![rle_inkyblackness only TEXT_MESSAGE_APP CDF](graphs/rle_inkyblackness_only_TEXT_MESSAGE_APP_cdf.png)

![rle_inkyblackness CDF](graphs/rle_inkyblackness_cdf.png)

### `flate_std`

![flate_std only TEXT_MESSAGE_APP CDF](graphs/flate_std_only_TEXT_MESSAGE_APP_cdf.png)

![flate_std CDF](graphs/flate_std_cdf.png)

### `lzw_std`

![lzw_std only TEXT_MESSAGE_APP CDF](graphs/lzw_std_only_TEXT_MESSAGE_APP_cdf.png)

![lzw_std CDF](graphs/lzw_std_cdf.png)

### `shoco_TextEn_tmthrgd`

![shoco_TextEn_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_TextEn_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_TextEn_tmthrgd CDF](graphs/shoco_TextEn_tmthrgd_cdf.png)

### `shoco_WordsEn_tmthrgd`

![shoco_WordsEn_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_WordsEn_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_WordsEn_tmthrgd CDF](graphs/shoco_WordsEn_tmthrgd_cdf.png)

### `shoco_Emails_tmthrgd`

![shoco_Emails_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_Emails_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_Emails_tmthrgd CDF](graphs/shoco_Emails_tmthrgd_cdf.png)

### `shoco_FilePath_tmthrgd`

![shoco_FilePath_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_FilePath_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_FilePath_tmthrgd CDF](graphs/shoco_FilePath_tmthrgd_cdf.png)

### `zstd_klauspost`

![zstd_klauspost only TEXT_MESSAGE_APP CDF](graphs/zstd_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

![zstd_klauspost CDF](graphs/zstd_klauspost_cdf.png)

### `zlib_klauspost`

![zlib_klauspost only TEXT_MESSAGE_APP CDF](graphs/zlib_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

![zlib_klauspost CDF](graphs/zlib_klauspost_cdf.png)

### `zlib_std`

![zlib_std only TEXT_MESSAGE_APP CDF](graphs/zlib_std_only_TEXT_MESSAGE_APP_cdf.png)

![zlib_std CDF](graphs/zlib_std_cdf.png)

### `smaz_cespare`

![smaz_cespare only TEXT_MESSAGE_APP CDF](graphs/smaz_cespare_only_TEXT_MESSAGE_APP_cdf.png)

![smaz_cespare CDF](graphs/smaz_cespare_cdf.png)

### `lz4_pierrec`

![lz4_pierrec only TEXT_MESSAGE_APP CDF](graphs/lz4_pierrec_only_TEXT_MESSAGE_APP_cdf.png)

![lz4_pierrec CDF](graphs/lz4_pierrec_cdf.png)

### `s2_klauspost`

![s2_klauspost only TEXT_MESSAGE_APP CDF](graphs/s2_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

![s2_klauspost CDF](graphs/s2_klauspost_cdf.png)

### `snappy_klauspost`

![snappy_klauspost only TEXT_MESSAGE_APP CDF](graphs/snappy_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

![snappy_klauspost CDF](graphs/snappy_klauspost_cdf.png)

### `gzip_klauspost`

![gzip_klauspost only TEXT_MESSAGE_APP CDF](graphs/gzip_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

![gzip_klauspost CDF](graphs/gzip_klauspost_cdf.png)

### `gzip_std`

![gzip_std only TEXT_MESSAGE_APP CDF](graphs/gzip_std_only_TEXT_MESSAGE_APP_cdf.png)

![gzip_std CDF](graphs/gzip_std_cdf.png)

