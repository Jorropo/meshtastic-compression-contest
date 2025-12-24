# Meshtastic Compression Showdown

This project contain benchmarks of various compression algorithms applied on a data of meshtastic packets.

For context a Reciprocal Compression Ratio **above** 1 means the compressed data is **bigger** than the uncompressed data.
One **bellow** 1 means the compressed data is **smaller** than the uncompressed data.

## Results

| Compressor | Average Reciprocal Compression Ratio | Average Reciprocal Compression Ratio (TEXT_MESSAGE_APP only) |
|------------|--------------------------------------|-------------------------------------------------------------|
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.9982 | 0.7383 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.9982 | 0.7380 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails` | 0.9982 | 0.7381 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9982 | 0.7379 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9982 | 0.7381 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails` | 0.9982 | 0.7384 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath` | 0.9982 | 0.7383 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.9982 | 0.7380 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9982 | 0.7379 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9982 | 0.7379 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed` | 0.9982 | 0.7382 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath` | 0.9982 | 0.7380 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9982 | 0.7381 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9982 | 0.7379 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9982 | 0.7381 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9982 | 0.7379 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.9982 | 0.7383 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9982 | 0.7381 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.9982 | 0.7380 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9982 | 0.7379 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails` | 0.9982 | 0.7381 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed` | 0.9982 | 0.7386 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed` | 0.9982 | 0.7382 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath` | 0.9982 | 0.7380 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.9982 | 0.7380 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9982 | 0.7379 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9982 | 0.7379 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9982 | 0.7385 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9982 | 0.7388 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed` | 0.9982 | 0.7387 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9982 | 0.7385 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn` | 0.9982 | 0.7388 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn` | 0.9982 | 0.7393 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_proposed` | 0.9982 | 0.7390 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed` | 0.9982 | 0.7386 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed` | 0.9982 | 0.7387 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed` | 0.9982 | 0.7388 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9982 | 0.7387 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9982 | 0.7387 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn` | 0.9982 | 0.7389 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9982 | 0.7385 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.9982 | 0.7388 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.9982 | 0.7386 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9982 | 0.7385 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath` | 0.9982 | 0.7387 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath` | 0.9982 | 0.7389 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed` | 0.9982 | 0.7386 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails` | 0.9982 | 0.7387 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9982 | 0.7385 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9982 | 0.7386 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.9982 | 0.7385 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath` | 0.9982 | 0.7386 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails` | 0.9982 | 0.7387 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails` | 0.9982 | 0.7390 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed` | 0.9983 | 0.7431 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9983 | 0.7426 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails` | 0.9983 | 0.7429 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9983 | 0.7426 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails` | 0.9983 | 0.7425 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed` | 0.9983 | 0.7426 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9983 | 0.7424 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.9983 | 0.7427 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.9983 | 0.7424 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9983 | 0.7424 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails` | 0.9983 | 0.7425 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9983 | 0.7424 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7425 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9983 | 0.7424 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath` | 0.9983 | 0.7424 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.9983 | 0.7424 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7423 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath` | 0.9983 | 0.7427 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9983 | 0.7423 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7428 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed` | 0.9983 | 0.7426 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7425 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7426 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9983 | 0.7426 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9983 | 0.7423 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath` | 0.9983 | 0.7424 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7423 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7456 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.9983 | 0.7432 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7431 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7433 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn` | 0.9983 | 0.7435 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed` | 0.9983 | 0.7432 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9983 | 0.7456 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed` | 0.9983 | 0.7460 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7458 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed` | 0.9983 | 0.7433 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7431 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed` | 0.9983 | 0.7432 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath` | 0.9983 | 0.7468 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9983 | 0.7431 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed` | 0.9983 | 0.7435 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails` | 0.9983 | 0.7458 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7455 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9983 | 0.7434 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath` | 0.9983 | 0.7436 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.9983 | 0.7431 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed` | 0.9983 | 0.7457 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails` | 0.9983 | 0.7465 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn` | 0.9983 | 0.7460 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9983 | 0.7456 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath` | 0.9983 | 0.7457 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails` | 0.9983 | 0.7457 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9983 | 0.7431 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed` | 0.9983 | 0.7467 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed` | 0.9983 | 0.7457 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9983 | 0.7456 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7464 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails` | 0.9983 | 0.7437 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails` | 0.9983 | 0.7433 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath` | 0.9983 | 0.7432 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails` | 0.9983 | 0.7433 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9983 | 0.7456 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9983 | 0.7432 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn` | 0.9983 | 0.7439 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9983 | 0.7434 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.9983 | 0.7434 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed` | 0.9983 | 0.7432 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_proposed` | 0.9983 | 0.7465 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9983 | 0.7431 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails` | 0.9983 | 0.7464 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath` | 0.9983 | 0.7433 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_proposed` | 0.9983 | 0.7458 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9983 | 0.7456 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn` | 0.9983 | 0.7434 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails` | 0.9983 | 0.7458 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath` | 0.9983 | 0.7458 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_proposed` | 0.9983 | 0.7436 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_proposed` | 0.9983 | 0.7475 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_Emails` | 0.9983 | 0.7475 |
| `sscc_Jorropo∪unishox2∪smaz` | 0.9983 | 0.7485 |
| `sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails` | 0.9983 | 0.7550 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9983 | 0.7538 |
| `sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails` | 0.9983 | 0.7549 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7540 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails` | 0.9983 | 0.7541 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath` | 0.9983 | 0.7540 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails` | 0.9983 | 0.7540 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath` | 0.9983 | 0.7540 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9983 | 0.7539 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed` | 0.9983 | 0.7539 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_proposed` | 0.9983 | 0.7541 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed` | 0.9983 | 0.7542 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7538 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9983 | 0.7538 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails` | 0.9983 | 0.7539 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9983 | 0.7538 |
| `sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_proposed` | 0.9983 | 0.7550 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed` | 0.9983 | 0.7539 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7538 |
| `sscc_Jorropo∪unishox2∪shoco_FilePath_proposed` | 0.9983 | 0.7553 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn` | 0.9983 | 0.7543 |
| `sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9983 | 0.7549 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9983 | 0.7538 |
| `sscc_Jorropo∪unishox2∪shoco_FilePath` | 0.9983 | 0.7553 |
| `sscc_Jorropo∪unishox2∪shoco_Emails` | 0.9983 | 0.7563 |
| `sscc_Jorropo∪unishox2∪shoco_Emails_proposed` | 0.9983 | 0.7562 |
| `sscc_Jorropo∪unishox2` | 0.9983 | 0.7578 |
| `unishox2_meshtastic` | 0.9983 | 0.7578 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9987 | 0.7599 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath` | 0.9987 | 0.7602 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9987 | 0.7600 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.9987 | 0.7600 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails` | 0.9987 | 0.7602 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath` | 0.9987 | 0.7601 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9987 | 0.7597 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed` | 0.9987 | 0.7604 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9987 | 0.7597 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9987 | 0.7598 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed` | 0.9987 | 0.7605 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.9987 | 0.7601 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails` | 0.9987 | 0.7602 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.9987 | 0.7601 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9987 | 0.7597 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9987 | 0.7599 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9987 | 0.7599 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.9987 | 0.7600 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9987 | 0.7603 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9987 | 0.7604 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9987 | 0.7605 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath` | 0.9987 | 0.7607 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.9987 | 0.7607 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails` | 0.9987 | 0.7608 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed` | 0.9987 | 0.7613 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.9987 | 0.7606 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9987 | 0.7602 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn` | 0.9987 | 0.7644 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9987 | 0.7635 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9987 | 0.7637 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9987 | 0.7636 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.9987 | 0.7639 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails` | 0.9987 | 0.7641 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed` | 0.9987 | 0.7639 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed` | 0.9987 | 0.7638 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath` | 0.9987 | 0.7640 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9987 | 0.7646 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed` | 0.9987 | 0.7649 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9987 | 0.7640 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed` | 0.9988 | 0.7645 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9988 | 0.7643 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9988 | 0.7652 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_proposed` | 0.9988 | 0.7655 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9988 | 0.7645 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9988 | 0.7648 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed` | 0.9988 | 0.7651 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn` | 0.9988 | 0.7661 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath` | 0.9988 | 0.7658 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails` | 0.9988 | 0.7658 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.9988 | 0.7656 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath` | 0.9988 | 0.7664 |
| `sscc_Jorropo∪smaz∪shoco_TextEn` | 0.9988 | 0.7669 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails` | 0.9988 | 0.7665 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.9988 | 0.7661 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9988 | 0.7703 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9988 | 0.7702 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9988 | 0.7701 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9988 | 0.7701 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9988 | 0.7703 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9988 | 0.7703 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9988 | 0.7701 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9988 | 0.7701 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath` | 0.9988 | 0.7706 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath` | 0.9988 | 0.7705 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.9988 | 0.7705 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.9988 | 0.7705 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails` | 0.9988 | 0.7705 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9988 | 0.7711 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9988 | 0.7710 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails` | 0.9988 | 0.7706 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9988 | 0.7708 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9988 | 0.7708 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.9988 | 0.7704 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.9988 | 0.7704 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed` | 0.9988 | 0.7709 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed` | 0.9988 | 0.7710 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath` | 0.9988 | 0.7714 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.9988 | 0.7713 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails` | 0.9988 | 0.7715 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.9988 | 0.7713 |
| `shoco_TextEn_proposed_tmthrgd_Jorropo` | 0.9988 | 0.7722 |
| `sscc_Jorropo∪shoco_TextEn_proposed` | 0.9988 | 0.7722 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9988 | 0.7768 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9988 | 0.7770 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.9988 | 0.7773 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9988 | 0.7770 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed` | 0.9988 | 0.7772 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed` | 0.9988 | 0.7772 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath` | 0.9988 | 0.7775 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails` | 0.9988 | 0.7776 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9988 | 0.7778 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9988 | 0.7787 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn` | 0.9988 | 0.7779 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9988 | 0.7786 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9988 | 0.7783 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed` | 0.9988 | 0.7790 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9988 | 0.7795 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed` | 0.9988 | 0.7785 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9988 | 0.7790 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed` | 0.9988 | 0.7794 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_Emails_proposed` | 0.9988 | 0.7799 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.9988 | 0.7806 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath` | 0.9988 | 0.7809 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.9988 | 0.7814 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails` | 0.9988 | 0.7809 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath` | 0.9988 | 0.7817 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn` | 0.9988 | 0.7813 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_Emails` | 0.9988 | 0.7818 |
| `sscc_Jorropo∪shoco_TextEn` | 0.9988 | 0.7825 |
| `shoco_TextEn_tmthrgd_Jorropo` | 0.9988 | 0.7825 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails` | 0.9989 | 0.7809 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed` | 0.9989 | 0.7807 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath` | 0.9989 | 0.7809 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9989 | 0.7804 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed` | 0.9989 | 0.7816 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed` | 0.9989 | 0.7806 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9989 | 0.7800 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9989 | 0.7801 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9989 | 0.7802 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9989 | 0.7808 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_proposed` | 0.9989 | 0.7820 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9989 | 0.7815 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9989 | 0.7812 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed` | 0.9989 | 0.7817 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails` | 0.9989 | 0.7833 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath` | 0.9989 | 0.7833 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn` | 0.9989 | 0.7841 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails` | 0.9989 | 0.7828 |
| `sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9989 | 0.7837 |
| `sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails_proposed` | 0.9989 | 0.7845 |
| `sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails` | 0.9989 | 0.7842 |
| `sscc_Jorropo∪smaz∪shoco_FilePath_proposed` | 0.9989 | 0.7857 |
| `sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails` | 0.9989 | 0.7859 |
| `sscc_Jorropo∪smaz∪shoco_Emails_proposed` | 0.9989 | 0.7871 |
| `sscc_Jorropo∪smaz∪shoco_FilePath` | 0.9989 | 0.7876 |
| `sscc_Jorropo∪smaz∪shoco_Emails` | 0.9989 | 0.7886 |
| `smaz_cespare_Jorropo` | 0.9990 | 0.7965 |
| `sscc_Jorropo∪smaz` | 0.9990 | 0.7965 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9991 | 0.8333 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.9991 | 0.8335 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.9991 | 0.8339 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.9991 | 0.8336 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed` | 0.9991 | 0.8344 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath` | 0.9991 | 0.8348 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9991 | 0.8372 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_proposed` | 0.9991 | 0.8370 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails` | 0.9991 | 0.8374 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.9991 | 0.8382 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.9991 | 0.8388 |
| `sscc_Jorropo∪shoco_WordsEn_proposed` | 0.9991 | 0.8388 |
| `shoco_WordsEn_proposed_tmthrgd_Jorropo` | 0.9991 | 0.8388 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed` | 0.9991 | 0.8391 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_Emails_proposed` | 0.9992 | 0.8423 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails` | 0.9992 | 0.8430 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath` | 0.9992 | 0.8439 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_Emails` | 0.9992 | 0.8467 |
| `sscc_Jorropo∪shoco_WordsEn` | 0.9992 | 0.8482 |
| `shoco_WordsEn_tmthrgd_Jorropo` | 0.9992 | 0.8482 |
| `sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.9993 | 0.8578 |
| `sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails` | 0.9993 | 0.8592 |
| `sscc_Jorropo∪shoco_FilePath∪shoco_Emails_proposed` | 0.9993 | 0.8602 |
| `sscc_Jorropo∪shoco_FilePath∪shoco_Emails` | 0.9993 | 0.8668 |
| `sscc_Jorropo∪shoco_FilePath_proposed` | 0.9993 | 0.8661 |
| `shoco_FilePath_proposed_tmthrgd_Jorropo` | 0.9993 | 0.8661 |
| `sscc_Jorropo∪shoco_Emails_proposed` | 0.9994 | 0.8755 |
| `shoco_Emails_proposed_tmthrgd_Jorropo` | 0.9994 | 0.8755 |
| `sscc_Jorropo∪shoco_FilePath` | 0.9994 | 0.8747 |
| `shoco_FilePath_tmthrgd_Jorropo` | 0.9994 | 0.8747 |
| `sscc_Jorropo∪shoco_Emails` | 0.9994 | 0.8838 |
| `shoco_Emails_tmthrgd_Jorropo` | 0.9994 | 0.8838 |
| `noop` | 1.0000 | 1.0000 |
| `lz4_cloudflareHC` | 1.0325 | 1.0429 |
| `lz4_cloudflare` | 1.0330 | 1.0434 |
| `flate_klauspost` | 1.0642 | 1.0595 |
| `rle_inkyblackness` | 1.0859 | 1.1225 |
| `flate_std` | 1.1264 | 1.1595 |
| `lzw_std` | 1.1473 | 1.1330 |
| `shoco_TextEn_proposed_tmthrgd` | 1.1715 | 0.8575 |
| `shoco_WordsEn_proposed_tmthrgd` | 1.1736 | 0.9241 |
| `shoco_Emails_proposed_tmthrgd` | 1.1755 | 0.9608 |
| `shoco_FilePath_proposed_tmthrgd` | 1.1758 | 0.9516 |
| `zstd_klauspost` | 1.1776 | 1.2531 |
| `zlib_klauspost` | 1.1852 | 1.2543 |
| `shoco_TextEn_tmthrgd` | 1.2200 | 0.8466 |
| `shoco_WordsEn_tmthrgd` | 1.2220 | 0.9133 |
| `shoco_Emails_tmthrgd` | 1.2240 | 0.9500 |
| `shoco_FilePath_tmthrgd` | 1.2243 | 0.9407 |
| `zlib_std` | 1.2475 | 1.3544 |
| `smaz_cespare` | 1.2514 | 0.8865 |
| `lz4_pierrec` | 1.2995 | 1.4766 |
| `s2_klauspost` | 1.3596 | 1.5730 |
| `snappy_klauspost` | 1.3613 | 1.5775 |
| `gzip_klauspost` | 1.4273 | 1.6440 |
| `gzip_std` | 1.4896 | 1.7440 |
## CDF Graphs

The following graphs show the cumulative distribution function (CDF) of the reciprocal compression ratios for each compressor.

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn`

![sscc_Jorropo∪unishox2∪shoco_TextEn CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz`

![sscc_Jorropo∪unishox2∪smaz CDF](graphs/sscc_Jorropo∪unishox2∪smaz_cdf.png)

![sscc_Jorropo∪unishox2∪smaz only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn`

![sscc_Jorropo∪unishox2∪shoco_WordsEn CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2`

![sscc_Jorropo∪unishox2 CDF](graphs/sscc_Jorropo∪unishox2_cdf.png)

![sscc_Jorropo∪unishox2 only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2_only_TEXT_MESSAGE_APP_cdf.png)

### `unishox2_meshtastic`

![unishox2_meshtastic CDF](graphs/unishox2_meshtastic_cdf.png)

![unishox2_meshtastic only TEXT_MESSAGE_APP CDF](graphs/unishox2_meshtastic_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn`

![sscc_Jorropo∪smaz∪shoco_TextEn CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_TextEn_proposed_tmthrgd_Jorropo`

![shoco_TextEn_proposed_tmthrgd_Jorropo CDF](graphs/shoco_TextEn_proposed_tmthrgd_Jorropo_cdf.png)

![shoco_TextEn_proposed_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_TextEn_proposed_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed`

![sscc_Jorropo∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn`

![sscc_Jorropo∪shoco_TextEn CDF](graphs/sscc_Jorropo∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_TextEn_tmthrgd_Jorropo`

![shoco_TextEn_tmthrgd_Jorropo CDF](graphs/shoco_TextEn_tmthrgd_Jorropo_cdf.png)

![shoco_TextEn_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_TextEn_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn`

![sscc_Jorropo∪smaz∪shoco_WordsEn CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `smaz_cespare_Jorropo`

![smaz_cespare_Jorropo CDF](graphs/smaz_cespare_Jorropo_cdf.png)

![smaz_cespare_Jorropo only TEXT_MESSAGE_APP CDF](graphs/smaz_cespare_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz`

![sscc_Jorropo∪smaz CDF](graphs/sscc_Jorropo∪smaz_cdf.png)

![sscc_Jorropo∪smaz only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_WordsEn_proposed_tmthrgd_Jorropo`

![shoco_WordsEn_proposed_tmthrgd_Jorropo CDF](graphs/shoco_WordsEn_proposed_tmthrgd_Jorropo_cdf.png)

![shoco_WordsEn_proposed_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_WordsEn_proposed_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath`

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn`

![sscc_Jorropo∪shoco_WordsEn CDF](graphs/sscc_Jorropo∪shoco_WordsEn_cdf.png)

![sscc_Jorropo∪shoco_WordsEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_WordsEn_tmthrgd_Jorropo`

![shoco_WordsEn_tmthrgd_Jorropo CDF](graphs/shoco_WordsEn_tmthrgd_Jorropo_cdf.png)

![shoco_WordsEn_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_WordsEn_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_FilePath_proposed_tmthrgd_Jorropo`

![shoco_FilePath_proposed_tmthrgd_Jorropo CDF](graphs/shoco_FilePath_proposed_tmthrgd_Jorropo_cdf.png)

![shoco_FilePath_proposed_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_FilePath_proposed_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_Emails_proposed_tmthrgd_Jorropo`

![shoco_Emails_proposed_tmthrgd_Jorropo CDF](graphs/shoco_Emails_proposed_tmthrgd_Jorropo_cdf.png)

![shoco_Emails_proposed_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_Emails_proposed_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_FilePath`

![sscc_Jorropo∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_FilePath_tmthrgd_Jorropo`

![shoco_FilePath_tmthrgd_Jorropo CDF](graphs/shoco_FilePath_tmthrgd_Jorropo_cdf.png)

![shoco_FilePath_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_FilePath_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_Emails`

![sscc_Jorropo∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_Emails_tmthrgd_Jorropo`

![shoco_Emails_tmthrgd_Jorropo CDF](graphs/shoco_Emails_tmthrgd_Jorropo_cdf.png)

![shoco_Emails_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_Emails_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

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

### `flate_std`

![flate_std CDF](graphs/flate_std_cdf.png)

![flate_std only TEXT_MESSAGE_APP CDF](graphs/flate_std_only_TEXT_MESSAGE_APP_cdf.png)

### `lzw_std`

![lzw_std CDF](graphs/lzw_std_cdf.png)

![lzw_std only TEXT_MESSAGE_APP CDF](graphs/lzw_std_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_TextEn_proposed_tmthrgd`

![shoco_TextEn_proposed_tmthrgd CDF](graphs/shoco_TextEn_proposed_tmthrgd_cdf.png)

![shoco_TextEn_proposed_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_TextEn_proposed_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_WordsEn_proposed_tmthrgd`

![shoco_WordsEn_proposed_tmthrgd CDF](graphs/shoco_WordsEn_proposed_tmthrgd_cdf.png)

![shoco_WordsEn_proposed_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_WordsEn_proposed_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_Emails_proposed_tmthrgd`

![shoco_Emails_proposed_tmthrgd CDF](graphs/shoco_Emails_proposed_tmthrgd_cdf.png)

![shoco_Emails_proposed_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_Emails_proposed_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_FilePath_proposed_tmthrgd`

![shoco_FilePath_proposed_tmthrgd CDF](graphs/shoco_FilePath_proposed_tmthrgd_cdf.png)

![shoco_FilePath_proposed_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_FilePath_proposed_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `zstd_klauspost`

![zstd_klauspost CDF](graphs/zstd_klauspost_cdf.png)

![zstd_klauspost only TEXT_MESSAGE_APP CDF](graphs/zstd_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

### `zlib_klauspost`

![zlib_klauspost CDF](graphs/zlib_klauspost_cdf.png)

![zlib_klauspost only TEXT_MESSAGE_APP CDF](graphs/zlib_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_TextEn_tmthrgd`

![shoco_TextEn_tmthrgd CDF](graphs/shoco_TextEn_tmthrgd_cdf.png)

![shoco_TextEn_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_TextEn_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_WordsEn_tmthrgd`

![shoco_WordsEn_tmthrgd CDF](graphs/shoco_WordsEn_tmthrgd_cdf.png)

![shoco_WordsEn_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_WordsEn_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_Emails_tmthrgd`

![shoco_Emails_tmthrgd CDF](graphs/shoco_Emails_tmthrgd_cdf.png)

![shoco_Emails_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_Emails_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_FilePath_tmthrgd`

![shoco_FilePath_tmthrgd CDF](graphs/shoco_FilePath_tmthrgd_cdf.png)

![shoco_FilePath_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_FilePath_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `zlib_std`

![zlib_std CDF](graphs/zlib_std_cdf.png)

![zlib_std only TEXT_MESSAGE_APP CDF](graphs/zlib_std_only_TEXT_MESSAGE_APP_cdf.png)

### `smaz_cespare`

![smaz_cespare CDF](graphs/smaz_cespare_cdf.png)

![smaz_cespare only TEXT_MESSAGE_APP CDF](graphs/smaz_cespare_only_TEXT_MESSAGE_APP_cdf.png)

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

