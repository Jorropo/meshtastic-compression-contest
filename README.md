# Meshtastic Compression Showdown

This project contain benchmarks of various compression algorithms applied on a data of meshtastic packets.

For context a Reciprocal Compression Ratio **above** 1 means the compressed data is **bigger** than the uncompressed data.
One **bellow** 1 means the compressed data is **smaller** than the uncompressed data.

## Results

| Compressor | Average Reciprocal Compression Ratio (TEXT_MESSAGE_APP only) | Average Reciprocal Compression Ratio |
|------------|--------------------------------------------------------------|--------------------------------------|
| `sscc_Jorropo_unishox2_smaz_shoco_WordsEn_shoco_TextEn` | 0.7382 | 0.9982 |
| `sscc_Jorropo_unishox2_smaz_shoco_TextEn` | 0.7386 | 0.9982 |
| `sscc_Jorropo_unishox2_shoco_WordsEn_shoco_TextEn` | 0.7426 | 0.9983 |
| `sscc_Jorropo_unishox2_shoco_TextEn` | 0.7431 | 0.9983 |
| `sscc_Jorropo_unishox2_smaz_shoco_WordsEn` | 0.7460 | 0.9983 |
| `sscc_Jorropo_unishox2_smaz` | 0.7485 | 0.9983 |
| `sscc_Jorropo_unishox2_shoco_WordsEn` | 0.7542 | 0.9983 |
| `unishox2_meshtastic` | 0.7578 | 0.9983 |
| `sscc_Jorropo_smaz_shoco_WordsEn_shoco_TextEn` | 0.7604 | 0.9987 |
| `sscc_Jorropo_smaz_shoco_TextEn` | 0.7613 | 0.9987 |
| `sscc_Jorropo_shoco_WordsEn_shoco_TextEn` | 0.7709 | 0.9988 |
| `shoco_TextEn_tmthrgd_Jorropo` | 0.7722 | 0.9988 |
| `sscc_Jorropo_smaz_shoco_WordsEn` | 0.7816 | 0.9989 |
| `smaz_cespare_Jorropo` | 0.7965 | 0.9990 |
| `shoco_WordsEn_tmthrgd_Jorropo` | 0.8388 | 0.9991 |
| `shoco_TextEn_tmthrgd` | 0.8575 | 1.1715 |
| `shoco_FilePath_tmthrgd_Jorropo` | 0.8661 | 0.9993 |
| `shoco_Emails_tmthrgd_Jorropo` | 0.8755 | 0.9994 |
| `smaz_cespare` | 0.8865 | 1.2514 |
| `shoco_WordsEn_tmthrgd` | 0.9241 | 1.1736 |
| `shoco_FilePath_tmthrgd` | 0.9516 | 1.1758 |
| `shoco_Emails_tmthrgd` | 0.9608 | 1.1755 |
| `noop` | 1.0000 | 1.0000 |
| `lz4_cloudflareHC` | 1.0429 | 1.0325 |
| `lz4_cloudflare` | 1.0434 | 1.0330 |
| `flate_klauspost` | 1.0595 | 1.0642 |
| `rle_inkyblackness` | 1.1225 | 1.0859 |
| `lzw_std` | 1.1330 | 1.1473 |
| `flate_std` | 1.1595 | 1.1264 |
| `zstd_klauspost` | 1.2531 | 1.1776 |
| `zlib_klauspost` | 1.2543 | 1.1852 |
| `zlib_std` | 1.3544 | 1.2475 |
| `lz4_pierrec` | 1.4766 | 1.2995 |
| `s2_klauspost` | 1.5730 | 1.3596 |
| `snappy_klauspost` | 1.5775 | 1.3613 |
| `gzip_klauspost` | 1.6440 | 1.4273 |
| `gzip_std` | 1.7440 | 1.4896 |
## CDF Graphs

The following graphs show the cumulative distribution function (CDF) of the reciprocal compression ratios for each compressor.

### `sscc_Jorropo_unishox2_smaz_shoco_WordsEn_shoco_TextEn`

![sscc_Jorropo_unishox2_smaz_shoco_WordsEn_shoco_TextEn CDF](graphs/sscc_Jorropo_unishox2_smaz_shoco_WordsEn_shoco_TextEn_cdf.png)

![sscc_Jorropo_unishox2_smaz_shoco_WordsEn_shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo_unishox2_smaz_shoco_WordsEn_shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo_unishox2_smaz_shoco_TextEn`

![sscc_Jorropo_unishox2_smaz_shoco_TextEn CDF](graphs/sscc_Jorropo_unishox2_smaz_shoco_TextEn_cdf.png)

![sscc_Jorropo_unishox2_smaz_shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo_unishox2_smaz_shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo_unishox2_shoco_WordsEn_shoco_TextEn`

![sscc_Jorropo_unishox2_shoco_WordsEn_shoco_TextEn CDF](graphs/sscc_Jorropo_unishox2_shoco_WordsEn_shoco_TextEn_cdf.png)

![sscc_Jorropo_unishox2_shoco_WordsEn_shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo_unishox2_shoco_WordsEn_shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo_unishox2_shoco_TextEn`

![sscc_Jorropo_unishox2_shoco_TextEn CDF](graphs/sscc_Jorropo_unishox2_shoco_TextEn_cdf.png)

![sscc_Jorropo_unishox2_shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo_unishox2_shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo_unishox2_smaz_shoco_WordsEn`

![sscc_Jorropo_unishox2_smaz_shoco_WordsEn CDF](graphs/sscc_Jorropo_unishox2_smaz_shoco_WordsEn_cdf.png)

![sscc_Jorropo_unishox2_smaz_shoco_WordsEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo_unishox2_smaz_shoco_WordsEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo_unishox2_smaz`

![sscc_Jorropo_unishox2_smaz CDF](graphs/sscc_Jorropo_unishox2_smaz_cdf.png)

![sscc_Jorropo_unishox2_smaz only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo_unishox2_smaz_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo_unishox2_shoco_WordsEn`

![sscc_Jorropo_unishox2_shoco_WordsEn CDF](graphs/sscc_Jorropo_unishox2_shoco_WordsEn_cdf.png)

![sscc_Jorropo_unishox2_shoco_WordsEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo_unishox2_shoco_WordsEn_only_TEXT_MESSAGE_APP_cdf.png)

### `unishox2_meshtastic`

![unishox2_meshtastic CDF](graphs/unishox2_meshtastic_cdf.png)

![unishox2_meshtastic only TEXT_MESSAGE_APP CDF](graphs/unishox2_meshtastic_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo_smaz_shoco_WordsEn_shoco_TextEn`

![sscc_Jorropo_smaz_shoco_WordsEn_shoco_TextEn CDF](graphs/sscc_Jorropo_smaz_shoco_WordsEn_shoco_TextEn_cdf.png)

![sscc_Jorropo_smaz_shoco_WordsEn_shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo_smaz_shoco_WordsEn_shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo_smaz_shoco_TextEn`

![sscc_Jorropo_smaz_shoco_TextEn CDF](graphs/sscc_Jorropo_smaz_shoco_TextEn_cdf.png)

![sscc_Jorropo_smaz_shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo_smaz_shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo_shoco_WordsEn_shoco_TextEn`

![sscc_Jorropo_shoco_WordsEn_shoco_TextEn CDF](graphs/sscc_Jorropo_shoco_WordsEn_shoco_TextEn_cdf.png)

![sscc_Jorropo_shoco_WordsEn_shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo_shoco_WordsEn_shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_TextEn_tmthrgd_Jorropo`

![shoco_TextEn_tmthrgd_Jorropo CDF](graphs/shoco_TextEn_tmthrgd_Jorropo_cdf.png)

![shoco_TextEn_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_TextEn_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo_smaz_shoco_WordsEn`

![sscc_Jorropo_smaz_shoco_WordsEn CDF](graphs/sscc_Jorropo_smaz_shoco_WordsEn_cdf.png)

![sscc_Jorropo_smaz_shoco_WordsEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo_smaz_shoco_WordsEn_only_TEXT_MESSAGE_APP_cdf.png)

### `smaz_cespare_Jorropo`

![smaz_cespare_Jorropo CDF](graphs/smaz_cespare_Jorropo_cdf.png)

![smaz_cespare_Jorropo only TEXT_MESSAGE_APP CDF](graphs/smaz_cespare_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_WordsEn_tmthrgd_Jorropo`

![shoco_WordsEn_tmthrgd_Jorropo CDF](graphs/shoco_WordsEn_tmthrgd_Jorropo_cdf.png)

![shoco_WordsEn_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_WordsEn_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_TextEn_tmthrgd`

![shoco_TextEn_tmthrgd CDF](graphs/shoco_TextEn_tmthrgd_cdf.png)

![shoco_TextEn_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_TextEn_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_FilePath_tmthrgd_Jorropo`

![shoco_FilePath_tmthrgd_Jorropo CDF](graphs/shoco_FilePath_tmthrgd_Jorropo_cdf.png)

![shoco_FilePath_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_FilePath_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_Emails_tmthrgd_Jorropo`

![shoco_Emails_tmthrgd_Jorropo CDF](graphs/shoco_Emails_tmthrgd_Jorropo_cdf.png)

![shoco_Emails_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_Emails_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `smaz_cespare`

![smaz_cespare CDF](graphs/smaz_cespare_cdf.png)

![smaz_cespare only TEXT_MESSAGE_APP CDF](graphs/smaz_cespare_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_WordsEn_tmthrgd`

![shoco_WordsEn_tmthrgd CDF](graphs/shoco_WordsEn_tmthrgd_cdf.png)

![shoco_WordsEn_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_WordsEn_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_FilePath_tmthrgd`

![shoco_FilePath_tmthrgd CDF](graphs/shoco_FilePath_tmthrgd_cdf.png)

![shoco_FilePath_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_FilePath_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_Emails_tmthrgd`

![shoco_Emails_tmthrgd CDF](graphs/shoco_Emails_tmthrgd_cdf.png)

![shoco_Emails_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_Emails_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `noop`

![noop CDF](graphs/noop_cdf.png)

![noop only TEXT_MESSAGE_APP CDF](graphs/noop_only_TEXT_MESSAGE_APP_cdf.png)

### `lz4_cloudflareHC`

![lz4_cloudflareHC CDF](graphs/lz4_cloudflareHC_cdf.png)

![lz4_cloudflareHC only TEXT_MESSAGE_APP CDF](graphs/lz4_cloudflareHC_only_TEXT_MESSAGE_APP_cdf.png)

### `lz4_cloudflare`

![lz4_cloudflare CDF](graphs/lz4_cloudflare_cdf.png)

![lz4_cloudflare only TEXT_MESSAGE_APP CDF](graphs/lz4_cloudflare_only_TEXT_MESSAGE_APP_cdf.png)

### `flate_klauspost`

![flate_klauspost CDF](graphs/flate_klauspost_cdf.png)

![flate_klauspost only TEXT_MESSAGE_APP CDF](graphs/flate_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

### `rle_inkyblackness`

![rle_inkyblackness CDF](graphs/rle_inkyblackness_cdf.png)

![rle_inkyblackness only TEXT_MESSAGE_APP CDF](graphs/rle_inkyblackness_only_TEXT_MESSAGE_APP_cdf.png)

### `lzw_std`

![lzw_std CDF](graphs/lzw_std_cdf.png)

![lzw_std only TEXT_MESSAGE_APP CDF](graphs/lzw_std_only_TEXT_MESSAGE_APP_cdf.png)

### `flate_std`

![flate_std CDF](graphs/flate_std_cdf.png)

![flate_std only TEXT_MESSAGE_APP CDF](graphs/flate_std_only_TEXT_MESSAGE_APP_cdf.png)

### `zstd_klauspost`

![zstd_klauspost CDF](graphs/zstd_klauspost_cdf.png)

![zstd_klauspost only TEXT_MESSAGE_APP CDF](graphs/zstd_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

### `zlib_klauspost`

![zlib_klauspost CDF](graphs/zlib_klauspost_cdf.png)

![zlib_klauspost only TEXT_MESSAGE_APP CDF](graphs/zlib_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

### `zlib_std`

![zlib_std CDF](graphs/zlib_std_cdf.png)

![zlib_std only TEXT_MESSAGE_APP CDF](graphs/zlib_std_only_TEXT_MESSAGE_APP_cdf.png)

### `lz4_pierrec`

![lz4_pierrec CDF](graphs/lz4_pierrec_cdf.png)

![lz4_pierrec only TEXT_MESSAGE_APP CDF](graphs/lz4_pierrec_only_TEXT_MESSAGE_APP_cdf.png)

### `s2_klauspost`

![s2_klauspost CDF](graphs/s2_klauspost_cdf.png)

![s2_klauspost only TEXT_MESSAGE_APP CDF](graphs/s2_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

### `snappy_klauspost`

![snappy_klauspost CDF](graphs/snappy_klauspost_cdf.png)

![snappy_klauspost only TEXT_MESSAGE_APP CDF](graphs/snappy_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

### `gzip_klauspost`

![gzip_klauspost CDF](graphs/gzip_klauspost_cdf.png)

![gzip_klauspost only TEXT_MESSAGE_APP CDF](graphs/gzip_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

### `gzip_std`

![gzip_std CDF](graphs/gzip_std_cdf.png)

![gzip_std only TEXT_MESSAGE_APP CDF](graphs/gzip_std_only_TEXT_MESSAGE_APP_cdf.png)

