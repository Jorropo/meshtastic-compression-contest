# Meshtastic Compression Showdown

This project contain benchmarks of various compression algorithms applied on a data of meshtastic packets.

For context a Reciprocal Compression Ratio **above** 1 means the compressed data is **bigger** than the uncompressed data.
One **bellow** 1 means the compressed data is **smaller** than the uncompressed data.

## Results

| Compressor | Average Reciprocal Compression Ratio (TEXT_MESSAGE_APP only) |
|------------|--------------------------------------------------------------|
| `unishox2_alpha_only` | 0.6976 |
| `meshtasticmodel_V1_EgonElbre` | 0.7154 |
| `meshtasticmodel_V4_EgonElbre` | 0.7154 |
| `meshtasticmodel_V5_EgonElbre` | 0.7154 |
| `meshtasticmodel_V6_EgonElbre` | 0.7154 |
| `meshtasticmodel_V10_EgonElbre` | 0.7158 |
| `meshtasticmodel_V8_EgonElbre` | 0.7158 |
| `meshtasticmodel_V9_EgonElbre` | 0.7158 |
| `meshtasticmodel_V7_EgonElbre` | 0.7173 |
| `meshtasticmodel_V3_EgonElbre` | 0.7188 |
| `snowflake_Jorropo` | 0.7196 |
| `unishox2_alpha_num_only` | 0.7325 |
| `unishox2_favor_alpha` | 0.7432 |
| `unishox2_no_uni_favor_text` | 0.7432 |
| `unishox2_no_uni` | 0.7515 |
| `unishox2_url` | 0.7519 |
| `unishox2_json_no_uni` | 0.7523 |
| `unishox2_alpha_num_sym_only` | 0.7534 |
| `unishox2_alpha_num_sym_only_text` | 0.7534 |
| `unishox2_default` | 0.7542 |
| `unishox2_xml` | 0.7542 |
| `unishox2_html` | 0.7549 |
| `unishox2_json` | 0.7549 |
| `unishox2_favor_dict` | 0.7568 |
| `shoco_TextEn_tmthrgd_Jorropo` | 0.7606 |
| `unishox2_favor_sym` | 0.7614 |
| `unishox2_favor_umlaut` | 0.7618 |
| `unishox2_no_dict` | 0.7622 |
| `smaz_cespare_Jorropo` | 0.7998 |
| `meshtasticmodel_V2_EgonElbre` | 0.8108 |
| `shoco_WordsEn_tmthrgd_Jorropo` | 0.8351 |
| `shoco_TextEn_tmthrgd` | 0.8419 |
| `shoco_FilePath_tmthrgd_Jorropo` | 0.8678 |
| `shoco_Emails_tmthrgd_Jorropo` | 0.8750 |
| `smaz_cespare` | 0.8807 |
| `shoco_WordsEn_tmthrgd` | 0.9160 |
| `meshtasticmodel_pbmodel-o1_EgonElbre` | 0.9328 |
| `meshtasticmodel_pbmodel-o2_EgonElbre` | 0.9328 |
| `meshtasticmodel_pbmodel-varint-o1_EgonElbre` | 0.9328 |
| `meshtasticmodel_pbmodel-varint-o2_EgonElbre` | 0.9328 |
| `meshtasticmodel_pbmodel-varint_EgonElbre` | 0.9328 |
| `meshtasticmodel_pbmodel_EgonElbre` | 0.9328 |
| `shoco_FilePath_tmthrgd` | 0.9487 |
| `shoco_Emails_tmthrgd` | 0.9559 |
| `noop` | 1.0000 |
| `lz4_cloudflareHC` | 1.0369 |
| `lz4_cloudflare` | 1.0376 |
| `flate_klauspost` | 1.0388 |
| `lzw_std` | 1.1071 |
| `rle_inkyblackness` | 1.1098 |
| `flate_std` | 1.1273 |
| `zlib_klauspost` | 1.2120 |
| `zlib_std` | 1.3005 |
| `lz4_pierrec` | 1.4229 |
| `s2_klauspost` | 1.5068 |
| `snappy_klauspost` | 1.5118 |
| `gzip_klauspost` | 1.5585 |
| `gzip_std` | 1.6470 |

| Compressor | Average Reciprocal Compression Ratio |
|------------|--------------------------------------|
| `snowflake_Jorropo` | 0.7837 |
| `meshtasticmodel_pbmodel-o1_EgonElbre` | 0.9541 |
| `meshtasticmodel_pbmodel-o2_EgonElbre` | 0.9541 |
| `meshtasticmodel_pbmodel_EgonElbre` | 0.9541 |
| `meshtasticmodel_pbmodel-varint-o1_EgonElbre` | 0.9541 |
| `meshtasticmodel_pbmodel-varint-o2_EgonElbre` | 0.9541 |
| `meshtasticmodel_pbmodel-varint_EgonElbre` | 0.9541 |
| `meshtasticmodel_V1_EgonElbre` | 0.9564 |
| `meshtasticmodel_V4_EgonElbre` | 0.9564 |
| `meshtasticmodel_V5_EgonElbre` | 0.9564 |
| `meshtasticmodel_V6_EgonElbre` | 0.9564 |
| `meshtasticmodel_V10_EgonElbre` | 0.9580 |
| `meshtasticmodel_V8_EgonElbre` | 0.9580 |
| `meshtasticmodel_V9_EgonElbre` | 0.9580 |
| `meshtasticmodel_V7_EgonElbre` | 0.9583 |
| `meshtasticmodel_V3_EgonElbre` | 0.9610 |
| `unishox2_alpha_only` | 0.9983 |
| `unishox2_alpha_num_only` | 0.9985 |
| `unishox2_favor_alpha` | 0.9986 |
| `unishox2_no_uni_favor_text` | 0.9986 |
| `unishox2_no_uni` | 0.9986 |
| `unishox2_url` | 0.9986 |
| `unishox2_json_no_uni` | 0.9986 |
| `unishox2_alpha_num_sym_only` | 0.9986 |
| `unishox2_alpha_num_sym_only_text` | 0.9986 |
| `unishox2_default` | 0.9986 |
| `unishox2_xml` | 0.9986 |
| `unishox2_html` | 0.9986 |
| `unishox2_json` | 0.9986 |
| `unishox2_favor_dict` | 0.9986 |
| `shoco_TextEn_tmthrgd_Jorropo` | 0.9987 |
| `unishox2_favor_sym` | 0.9987 |
| `unishox2_favor_umlaut` | 0.9987 |
| `unishox2_no_dict` | 0.9987 |
| `smaz_cespare_Jorropo` | 0.9989 |
| `shoco_WordsEn_tmthrgd_Jorropo` | 0.9991 |
| `shoco_FilePath_tmthrgd_Jorropo` | 0.9993 |
| `shoco_Emails_tmthrgd_Jorropo` | 0.9993 |
| `noop` | 1.0000 |
| `meshtasticmodel_V2_EgonElbre` | 1.0147 |
| `lz4_cloudflareHC` | 1.0340 |
| `lz4_cloudflare` | 1.0344 |
| `flate_klauspost` | 1.0660 |
| `rle_inkyblackness` | 1.0849 |
| `flate_std` | 1.1292 |
| `lzw_std` | 1.1484 |
| `shoco_TextEn_tmthrgd` | 1.1694 |
| `shoco_WordsEn_tmthrgd` | 1.1705 |
| `shoco_Emails_tmthrgd` | 1.1718 |
| `shoco_FilePath_tmthrgd` | 1.1728 |
| `zlib_klauspost` | 1.1903 |
| `smaz_cespare` | 1.2517 |
| `zlib_std` | 1.2535 |
| `lz4_pierrec` | 1.3083 |
| `s2_klauspost` | 1.3701 |
| `snappy_klauspost` | 1.3722 |
| `gzip_klauspost` | 1.4390 |
| `gzip_std` | 1.5022 |

## CDF Graphs

The following graphs show the cumulative distribution function (CDF) of the reciprocal compression ratios for each compressor.

### `snowflake_Jorropo`

![snowflake_Jorropo only TEXT_MESSAGE_APP CDF](graphs/snowflake_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![snowflake_Jorropo CDF](graphs/snowflake_Jorropo_cdf.png)

### `meshtasticmodel_pbmodel-o1_EgonElbre`

![meshtasticmodel_pbmodel-o1_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_pbmodel-o1_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_pbmodel-o1_EgonElbre CDF](graphs/meshtasticmodel_pbmodel-o1_EgonElbre_cdf.png)

### `meshtasticmodel_pbmodel-o2_EgonElbre`

![meshtasticmodel_pbmodel-o2_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_pbmodel-o2_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_pbmodel-o2_EgonElbre CDF](graphs/meshtasticmodel_pbmodel-o2_EgonElbre_cdf.png)

### `meshtasticmodel_pbmodel_EgonElbre`

![meshtasticmodel_pbmodel_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_pbmodel_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_pbmodel_EgonElbre CDF](graphs/meshtasticmodel_pbmodel_EgonElbre_cdf.png)

### `meshtasticmodel_pbmodel-varint-o1_EgonElbre`

![meshtasticmodel_pbmodel-varint-o1_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_pbmodel-varint-o1_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_pbmodel-varint-o1_EgonElbre CDF](graphs/meshtasticmodel_pbmodel-varint-o1_EgonElbre_cdf.png)

### `meshtasticmodel_pbmodel-varint-o2_EgonElbre`

![meshtasticmodel_pbmodel-varint-o2_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_pbmodel-varint-o2_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_pbmodel-varint-o2_EgonElbre CDF](graphs/meshtasticmodel_pbmodel-varint-o2_EgonElbre_cdf.png)

### `meshtasticmodel_pbmodel-varint_EgonElbre`

![meshtasticmodel_pbmodel-varint_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_pbmodel-varint_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_pbmodel-varint_EgonElbre CDF](graphs/meshtasticmodel_pbmodel-varint_EgonElbre_cdf.png)

### `meshtasticmodel_V1_EgonElbre`

![meshtasticmodel_V1_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V1_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V1_EgonElbre CDF](graphs/meshtasticmodel_V1_EgonElbre_cdf.png)

### `meshtasticmodel_V4_EgonElbre`

![meshtasticmodel_V4_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V4_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V4_EgonElbre CDF](graphs/meshtasticmodel_V4_EgonElbre_cdf.png)

### `meshtasticmodel_V5_EgonElbre`

![meshtasticmodel_V5_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V5_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V5_EgonElbre CDF](graphs/meshtasticmodel_V5_EgonElbre_cdf.png)

### `meshtasticmodel_V6_EgonElbre`

![meshtasticmodel_V6_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V6_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V6_EgonElbre CDF](graphs/meshtasticmodel_V6_EgonElbre_cdf.png)

### `meshtasticmodel_V10_EgonElbre`

![meshtasticmodel_V10_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V10_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V10_EgonElbre CDF](graphs/meshtasticmodel_V10_EgonElbre_cdf.png)

### `meshtasticmodel_V8_EgonElbre`

![meshtasticmodel_V8_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V8_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V8_EgonElbre CDF](graphs/meshtasticmodel_V8_EgonElbre_cdf.png)

### `meshtasticmodel_V9_EgonElbre`

![meshtasticmodel_V9_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V9_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V9_EgonElbre CDF](graphs/meshtasticmodel_V9_EgonElbre_cdf.png)

### `meshtasticmodel_V7_EgonElbre`

![meshtasticmodel_V7_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V7_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V7_EgonElbre CDF](graphs/meshtasticmodel_V7_EgonElbre_cdf.png)

### `meshtasticmodel_V3_EgonElbre`

![meshtasticmodel_V3_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V3_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V3_EgonElbre CDF](graphs/meshtasticmodel_V3_EgonElbre_cdf.png)

### `unishox2_alpha_only`

![unishox2_alpha_only only TEXT_MESSAGE_APP CDF](graphs/unishox2_alpha_only_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_alpha_only CDF](graphs/unishox2_alpha_only_cdf.png)

### `unishox2_alpha_num_only`

![unishox2_alpha_num_only only TEXT_MESSAGE_APP CDF](graphs/unishox2_alpha_num_only_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_alpha_num_only CDF](graphs/unishox2_alpha_num_only_cdf.png)

### `unishox2_favor_alpha`

![unishox2_favor_alpha only TEXT_MESSAGE_APP CDF](graphs/unishox2_favor_alpha_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_favor_alpha CDF](graphs/unishox2_favor_alpha_cdf.png)

### `unishox2_no_uni_favor_text`

![unishox2_no_uni_favor_text only TEXT_MESSAGE_APP CDF](graphs/unishox2_no_uni_favor_text_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_no_uni_favor_text CDF](graphs/unishox2_no_uni_favor_text_cdf.png)

### `unishox2_no_uni`

![unishox2_no_uni only TEXT_MESSAGE_APP CDF](graphs/unishox2_no_uni_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_no_uni CDF](graphs/unishox2_no_uni_cdf.png)

### `unishox2_url`

![unishox2_url only TEXT_MESSAGE_APP CDF](graphs/unishox2_url_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_url CDF](graphs/unishox2_url_cdf.png)

### `unishox2_json_no_uni`

![unishox2_json_no_uni only TEXT_MESSAGE_APP CDF](graphs/unishox2_json_no_uni_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_json_no_uni CDF](graphs/unishox2_json_no_uni_cdf.png)

### `unishox2_alpha_num_sym_only`

![unishox2_alpha_num_sym_only only TEXT_MESSAGE_APP CDF](graphs/unishox2_alpha_num_sym_only_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_alpha_num_sym_only CDF](graphs/unishox2_alpha_num_sym_only_cdf.png)

### `unishox2_alpha_num_sym_only_text`

![unishox2_alpha_num_sym_only_text only TEXT_MESSAGE_APP CDF](graphs/unishox2_alpha_num_sym_only_text_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_alpha_num_sym_only_text CDF](graphs/unishox2_alpha_num_sym_only_text_cdf.png)

### `unishox2_default`

![unishox2_default only TEXT_MESSAGE_APP CDF](graphs/unishox2_default_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_default CDF](graphs/unishox2_default_cdf.png)

### `unishox2_xml`

![unishox2_xml only TEXT_MESSAGE_APP CDF](graphs/unishox2_xml_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_xml CDF](graphs/unishox2_xml_cdf.png)

### `unishox2_html`

![unishox2_html only TEXT_MESSAGE_APP CDF](graphs/unishox2_html_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_html CDF](graphs/unishox2_html_cdf.png)

### `unishox2_json`

![unishox2_json only TEXT_MESSAGE_APP CDF](graphs/unishox2_json_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_json CDF](graphs/unishox2_json_cdf.png)

### `unishox2_favor_dict`

![unishox2_favor_dict only TEXT_MESSAGE_APP CDF](graphs/unishox2_favor_dict_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_favor_dict CDF](graphs/unishox2_favor_dict_cdf.png)

### `shoco_TextEn_tmthrgd_Jorropo`

![shoco_TextEn_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_TextEn_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_TextEn_tmthrgd_Jorropo CDF](graphs/shoco_TextEn_tmthrgd_Jorropo_cdf.png)

### `unishox2_favor_sym`

![unishox2_favor_sym only TEXT_MESSAGE_APP CDF](graphs/unishox2_favor_sym_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_favor_sym CDF](graphs/unishox2_favor_sym_cdf.png)

### `unishox2_favor_umlaut`

![unishox2_favor_umlaut only TEXT_MESSAGE_APP CDF](graphs/unishox2_favor_umlaut_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_favor_umlaut CDF](graphs/unishox2_favor_umlaut_cdf.png)

### `unishox2_no_dict`

![unishox2_no_dict only TEXT_MESSAGE_APP CDF](graphs/unishox2_no_dict_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_no_dict CDF](graphs/unishox2_no_dict_cdf.png)

### `smaz_cespare_Jorropo`

![smaz_cespare_Jorropo only TEXT_MESSAGE_APP CDF](graphs/smaz_cespare_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![smaz_cespare_Jorropo CDF](graphs/smaz_cespare_Jorropo_cdf.png)

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

### `meshtasticmodel_V2_EgonElbre`

![meshtasticmodel_V2_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V2_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V2_EgonElbre CDF](graphs/meshtasticmodel_V2_EgonElbre_cdf.png)

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

### `zlib_klauspost`

![zlib_klauspost only TEXT_MESSAGE_APP CDF](graphs/zlib_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

![zlib_klauspost CDF](graphs/zlib_klauspost_cdf.png)

### `smaz_cespare`

![smaz_cespare only TEXT_MESSAGE_APP CDF](graphs/smaz_cespare_only_TEXT_MESSAGE_APP_cdf.png)

![smaz_cespare CDF](graphs/smaz_cespare_cdf.png)

### `zlib_std`

![zlib_std only TEXT_MESSAGE_APP CDF](graphs/zlib_std_only_TEXT_MESSAGE_APP_cdf.png)

![zlib_std CDF](graphs/zlib_std_cdf.png)

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

