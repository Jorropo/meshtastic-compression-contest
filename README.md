# Meshtastic Compression Showdown

This project contain benchmarks of various compression algorithms applied on a data of meshtastic packets.

For context a Reciprocal Compression Ratio **above** 1 means the compressed data is **bigger** than the uncompressed data.
One **bellow** 1 means the compressed data is **smaller** than the uncompressed data.

## Results

| Compressor | Average Reciprocal Compression Ratio (TEXT_MESSAGE_APP only) | Average Reciprocal Compression Ratio |
|------------|--------------------------------------------------------------|--------------------------------------|
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7379 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.7379 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7379 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.7379 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.7379 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.7379 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.7379 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.7379 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath` | 0.7380 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.7380 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.7380 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath` | 0.7380 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.7380 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.7380 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails` | 0.7381 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails` | 0.7381 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7381 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.7381 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.7381 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.7381 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed` | 0.7382 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed` | 0.7382 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.7383 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath` | 0.7383 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.7383 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails` | 0.7384 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7385 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7385 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.7385 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.7385 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.7385 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.7385 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.7386 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed` | 0.7386 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed` | 0.7386 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed` | 0.7386 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.7386 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath` | 0.7386 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed` | 0.7387 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails` | 0.7387 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath` | 0.7387 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed` | 0.7387 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7387 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.7387 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails` | 0.7387 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.7388 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn` | 0.7388 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.7388 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed` | 0.7388 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn` | 0.7389 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath` | 0.7389 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_proposed` | 0.7390 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails` | 0.7390 | 0.9982 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn` | 0.7393 | 0.9982 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7423 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.7423 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.7423 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7423 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.7424 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.7424 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.7424 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.7424 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.7424 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath` | 0.7424 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath` | 0.7424 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.7424 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.7425 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.7425 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails` | 0.7425 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails` | 0.7425 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7426 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.7426 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.7426 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.7426 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed` | 0.7426 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed` | 0.7426 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.7427 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath` | 0.7427 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.7428 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails` | 0.7429 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7431 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7431 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.7431 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.7431 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed` | 0.7431 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.7431 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.7431 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.7432 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.7432 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed` | 0.7432 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed` | 0.7432 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath` | 0.7432 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed` | 0.7432 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails` | 0.7433 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath` | 0.7433 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed` | 0.7433 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails` | 0.7433 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7433 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.7434 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.7434 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn` | 0.7434 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.7434 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn` | 0.7435 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed` | 0.7435 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath` | 0.7436 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_proposed` | 0.7436 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails` | 0.7437 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_TextEn` | 0.7439 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7455 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7456 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.7456 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.7456 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.7456 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.7456 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.7456 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails` | 0.7457 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed` | 0.7457 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed` | 0.7457 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath` | 0.7457 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed` | 0.7458 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath` | 0.7458 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails` | 0.7458 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_proposed` | 0.7458 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails` | 0.7458 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed` | 0.7460 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn` | 0.7460 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7464 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails` | 0.7464 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_proposed` | 0.7465 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails` | 0.7465 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed` | 0.7467 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath` | 0.7468 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_proposed` | 0.7475 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz∪shoco_Emails` | 0.7475 | 0.9983 |
| `sscc_Jorropo∪unishox2∪smaz` | 0.7485 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7538 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7538 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.7538 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.7538 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.7538 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.7538 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.7539 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails` | 0.7539 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed` | 0.7539 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed` | 0.7539 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath` | 0.7540 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_proposed` | 0.7540 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath` | 0.7540 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails` | 0.7540 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_proposed` | 0.7541 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails` | 0.7541 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed` | 0.7542 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_WordsEn` | 0.7543 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7549 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails` | 0.7549 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_proposed` | 0.7550 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails` | 0.7550 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_FilePath_proposed` | 0.7553 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_FilePath` | 0.7553 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_Emails_proposed` | 0.7562 | 0.9983 |
| `sscc_Jorropo∪unishox2∪shoco_Emails` | 0.7563 | 0.9983 |
| `unishox2_meshtastic` | 0.7578 | 0.9983 |
| `sscc_Jorropo∪unishox2` | 0.7578 | 0.9983 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7597 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.7597 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7597 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.7598 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.7599 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.7599 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.7599 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.7600 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.7600 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.7600 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.7601 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath` | 0.7601 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.7601 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath` | 0.7602 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails` | 0.7602 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails` | 0.7602 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7602 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.7603 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.7604 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed` | 0.7604 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.7605 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed` | 0.7605 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.7606 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.7607 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath` | 0.7607 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails` | 0.7608 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_TextEn_proposed` | 0.7613 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7635 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.7636 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.7637 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed` | 0.7638 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed` | 0.7639 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.7639 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7640 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath` | 0.7640 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails` | 0.7641 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.7643 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn` | 0.7644 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed` | 0.7645 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7645 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.7646 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.7648 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed` | 0.7649 | 0.9987 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed` | 0.7651 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.7652 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_proposed` | 0.7655 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.7656 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath` | 0.7658 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails` | 0.7658 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.7661 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn` | 0.7661 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath` | 0.7664 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails` | 0.7665 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_TextEn` | 0.7669 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7701 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.7701 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7701 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.7701 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.7702 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.7703 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.7703 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.7703 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.7704 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.7704 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.7705 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath` | 0.7705 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.7705 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails` | 0.7705 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath` | 0.7706 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails` | 0.7706 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7708 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.7708 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed` | 0.7709 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed` | 0.7710 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.7710 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.7711 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_proposed` | 0.7713 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed` | 0.7713 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath` | 0.7714 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails` | 0.7715 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn_proposed` | 0.7722 | 0.9988 |
| `shoco_TextEn_proposed_tmthrgd_Jorropo` | 0.7722 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7768 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.7770 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.7770 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed` | 0.7772 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed` | 0.7772 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.7773 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath` | 0.7775 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails` | 0.7776 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7778 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn` | 0.7779 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.7783 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed` | 0.7785 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7786 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.7787 | 0.9988 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed` | 0.7790 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.7790 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed` | 0.7794 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.7795 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_Emails_proposed` | 0.7799 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7800 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.7801 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.7802 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.7804 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed` | 0.7806 | 0.9989 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.7806 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed` | 0.7807 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7808 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath` | 0.7809 | 0.9989 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath` | 0.7809 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails` | 0.7809 | 0.9989 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails` | 0.7809 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.7812 | 0.9989 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn` | 0.7813 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails` | 0.7814 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.7815 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed` | 0.7816 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed` | 0.7817 | 0.9989 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath` | 0.7817 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn∪shoco_Emails` | 0.7818 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_proposed` | 0.7820 | 0.9989 |
| `shoco_TextEn_tmthrgd_Jorropo` | 0.7825 | 0.9988 |
| `sscc_Jorropo∪shoco_TextEn` | 0.7825 | 0.9988 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails` | 0.7828 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath` | 0.7833 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails` | 0.7833 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.7837 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_WordsEn` | 0.7841 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails` | 0.7842 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails_proposed` | 0.7845 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_FilePath_proposed` | 0.7857 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails` | 0.7859 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_Emails_proposed` | 0.7871 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_FilePath` | 0.7876 | 0.9989 |
| `sscc_Jorropo∪smaz∪shoco_Emails` | 0.7886 | 0.9989 |
| `sscc_Jorropo∪smaz` | 0.7965 | 0.9990 |
| `smaz_cespare_Jorropo` | 0.7965 | 0.9990 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.8333 | 0.9991 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed` | 0.8335 | 0.9991 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails` | 0.8336 | 0.9991 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails` | 0.8339 | 0.9991 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed` | 0.8344 | 0.9991 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath` | 0.8348 | 0.9991 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_proposed` | 0.8370 | 0.9991 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.8372 | 0.9991 |
| `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails` | 0.8374 | 0.9991 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails` | 0.8382 | 0.9991 |
| `sscc_Jorropo∪shoco_WordsEn_proposed` | 0.8388 | 0.9991 |
| `shoco_WordsEn_proposed_tmthrgd_Jorropo` | 0.8388 | 0.9991 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed` | 0.8388 | 0.9991 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed` | 0.8391 | 0.9991 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_Emails_proposed` | 0.8423 | 0.9992 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails` | 0.8430 | 0.9992 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath` | 0.8439 | 0.9992 |
| `shoco_TextEn_tmthrgd` | 0.8466 | 1.2200 |
| `sscc_Jorropo∪shoco_WordsEn∪shoco_Emails` | 0.8467 | 0.9992 |
| `sscc_Jorropo∪shoco_WordsEn` | 0.8482 | 0.9992 |
| `shoco_WordsEn_tmthrgd_Jorropo` | 0.8482 | 0.9992 |
| `shoco_TextEn_proposed_tmthrgd` | 0.8575 | 1.1715 |
| `sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_proposed` | 0.8578 | 0.9993 |
| `sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails` | 0.8592 | 0.9993 |
| `sscc_Jorropo∪shoco_FilePath∪shoco_Emails_proposed` | 0.8602 | 0.9993 |
| `sscc_Jorropo∪shoco_FilePath_proposed` | 0.8661 | 0.9993 |
| `shoco_FilePath_proposed_tmthrgd_Jorropo` | 0.8661 | 0.9993 |
| `sscc_Jorropo∪shoco_FilePath∪shoco_Emails` | 0.8668 | 0.9993 |
| `sscc_Jorropo∪shoco_FilePath` | 0.8747 | 0.9994 |
| `shoco_FilePath_tmthrgd_Jorropo` | 0.8747 | 0.9994 |
| `shoco_Emails_proposed_tmthrgd_Jorropo` | 0.8755 | 0.9994 |
| `sscc_Jorropo∪shoco_Emails_proposed` | 0.8755 | 0.9994 |
| `sscc_Jorropo∪shoco_Emails` | 0.8838 | 0.9994 |
| `shoco_Emails_tmthrgd_Jorropo` | 0.8838 | 0.9994 |
| `smaz_cespare` | 0.8865 | 1.2514 |
| `shoco_WordsEn_tmthrgd` | 0.9133 | 1.2220 |
| `shoco_WordsEn_proposed_tmthrgd` | 0.9241 | 1.1736 |
| `shoco_FilePath_tmthrgd` | 0.9407 | 1.2243 |
| `shoco_Emails_tmthrgd` | 0.9500 | 1.2240 |
| `shoco_FilePath_proposed_tmthrgd` | 0.9516 | 1.1758 |
| `shoco_Emails_proposed_tmthrgd` | 0.9608 | 1.1755 |
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

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn`

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_TextEn`

![sscc_Jorropo∪unishox2∪shoco_TextEn CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn`

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_WordsEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz∪shoco_Emails`

![sscc_Jorropo∪unishox2∪smaz∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪smaz∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪smaz`

![sscc_Jorropo∪unishox2∪smaz CDF](graphs/sscc_Jorropo∪unishox2∪smaz_cdf.png)

![sscc_Jorropo∪unishox2∪smaz only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪smaz_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed`

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_WordsEn`

![sscc_Jorropo∪unishox2∪shoco_WordsEn CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_WordsEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_WordsEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_FilePath_proposed`

![sscc_Jorropo∪unishox2∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_FilePath`

![sscc_Jorropo∪unishox2∪shoco_FilePath CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_Emails_proposed`

![sscc_Jorropo∪unishox2∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪unishox2∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2∪shoco_Emails`

![sscc_Jorropo∪unishox2∪shoco_Emails CDF](graphs/sscc_Jorropo∪unishox2∪shoco_Emails_cdf.png)

![sscc_Jorropo∪unishox2∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `unishox2_meshtastic`

![unishox2_meshtastic CDF](graphs/unishox2_meshtastic_cdf.png)

![unishox2_meshtastic only TEXT_MESSAGE_APP CDF](graphs/unishox2_meshtastic_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪unishox2`

![sscc_Jorropo∪unishox2 CDF](graphs/sscc_Jorropo∪unishox2_cdf.png)

![sscc_Jorropo∪unishox2 only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪unishox2_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_TextEn`

![sscc_Jorropo∪smaz∪shoco_TextEn CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪smaz∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn_proposed`

![sscc_Jorropo∪shoco_TextEn_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_TextEn_proposed_tmthrgd_Jorropo`

![shoco_TextEn_proposed_tmthrgd_Jorropo CDF](graphs/shoco_TextEn_proposed_tmthrgd_Jorropo_cdf.png)

![shoco_TextEn_proposed_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_TextEn_proposed_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_TextEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn`

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_FilePath`

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn∪shoco_Emails`

![sscc_Jorropo∪shoco_TextEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_TextEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_TextEn_tmthrgd_Jorropo`

![shoco_TextEn_tmthrgd_Jorropo CDF](graphs/shoco_TextEn_tmthrgd_Jorropo_cdf.png)

![shoco_TextEn_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_TextEn_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_TextEn`

![sscc_Jorropo∪shoco_TextEn CDF](graphs/sscc_Jorropo∪shoco_TextEn_cdf.png)

![sscc_Jorropo∪shoco_TextEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_TextEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_WordsEn`

![sscc_Jorropo∪smaz∪shoco_WordsEn CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_cdf.png)

![sscc_Jorropo∪smaz∪shoco_WordsEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_WordsEn_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

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

### `sscc_Jorropo∪smaz`

![sscc_Jorropo∪smaz CDF](graphs/sscc_Jorropo∪smaz_cdf.png)

![sscc_Jorropo∪smaz only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪smaz_only_TEXT_MESSAGE_APP_cdf.png)

### `smaz_cespare_Jorropo`

![smaz_cespare_Jorropo CDF](graphs/smaz_cespare_Jorropo_cdf.png)

![smaz_cespare_Jorropo only TEXT_MESSAGE_APP CDF](graphs/smaz_cespare_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn_proposed`

![sscc_Jorropo∪shoco_WordsEn_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_WordsEn_proposed_tmthrgd_Jorropo`

![shoco_WordsEn_proposed_tmthrgd_Jorropo CDF](graphs/shoco_WordsEn_proposed_tmthrgd_Jorropo_cdf.png)

![shoco_WordsEn_proposed_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_WordsEn_proposed_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

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

### `shoco_TextEn_tmthrgd`

![shoco_TextEn_tmthrgd CDF](graphs/shoco_TextEn_tmthrgd_cdf.png)

![shoco_TextEn_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_TextEn_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn∪shoco_Emails`

![sscc_Jorropo∪shoco_WordsEn∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_WordsEn∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_WordsEn`

![sscc_Jorropo∪shoco_WordsEn CDF](graphs/sscc_Jorropo∪shoco_WordsEn_cdf.png)

![sscc_Jorropo∪shoco_WordsEn only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_WordsEn_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_WordsEn_tmthrgd_Jorropo`

![shoco_WordsEn_tmthrgd_Jorropo CDF](graphs/shoco_WordsEn_tmthrgd_Jorropo_cdf.png)

![shoco_WordsEn_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_WordsEn_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_TextEn_proposed_tmthrgd`

![shoco_TextEn_proposed_tmthrgd CDF](graphs/shoco_TextEn_proposed_tmthrgd_cdf.png)

![shoco_TextEn_proposed_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_TextEn_proposed_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails`

![sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_FilePath_proposed∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_FilePath∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_FilePath∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_FilePath∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_FilePath∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_FilePath∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_FilePath_proposed`

![sscc_Jorropo∪shoco_FilePath_proposed CDF](graphs/sscc_Jorropo∪shoco_FilePath_proposed_cdf.png)

![sscc_Jorropo∪shoco_FilePath_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_FilePath_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_FilePath_proposed_tmthrgd_Jorropo`

![shoco_FilePath_proposed_tmthrgd_Jorropo CDF](graphs/shoco_FilePath_proposed_tmthrgd_Jorropo_cdf.png)

![shoco_FilePath_proposed_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_FilePath_proposed_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_FilePath∪shoco_Emails`

![sscc_Jorropo∪shoco_FilePath∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_FilePath∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_FilePath∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_FilePath∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_FilePath`

![sscc_Jorropo∪shoco_FilePath CDF](graphs/sscc_Jorropo∪shoco_FilePath_cdf.png)

![sscc_Jorropo∪shoco_FilePath only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_FilePath_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_FilePath_tmthrgd_Jorropo`

![shoco_FilePath_tmthrgd_Jorropo CDF](graphs/shoco_FilePath_tmthrgd_Jorropo_cdf.png)

![shoco_FilePath_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_FilePath_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_Emails_proposed_tmthrgd_Jorropo`

![shoco_Emails_proposed_tmthrgd_Jorropo CDF](graphs/shoco_Emails_proposed_tmthrgd_Jorropo_cdf.png)

![shoco_Emails_proposed_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_Emails_proposed_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_Emails_proposed`

![sscc_Jorropo∪shoco_Emails_proposed CDF](graphs/sscc_Jorropo∪shoco_Emails_proposed_cdf.png)

![sscc_Jorropo∪shoco_Emails_proposed only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_Emails_proposed_only_TEXT_MESSAGE_APP_cdf.png)

### `sscc_Jorropo∪shoco_Emails`

![sscc_Jorropo∪shoco_Emails CDF](graphs/sscc_Jorropo∪shoco_Emails_cdf.png)

![sscc_Jorropo∪shoco_Emails only TEXT_MESSAGE_APP CDF](graphs/sscc_Jorropo∪shoco_Emails_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_Emails_tmthrgd_Jorropo`

![shoco_Emails_tmthrgd_Jorropo CDF](graphs/shoco_Emails_tmthrgd_Jorropo_cdf.png)

![shoco_Emails_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_Emails_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

### `smaz_cespare`

![smaz_cespare CDF](graphs/smaz_cespare_cdf.png)

![smaz_cespare only TEXT_MESSAGE_APP CDF](graphs/smaz_cespare_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_WordsEn_tmthrgd`

![shoco_WordsEn_tmthrgd CDF](graphs/shoco_WordsEn_tmthrgd_cdf.png)

![shoco_WordsEn_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_WordsEn_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_WordsEn_proposed_tmthrgd`

![shoco_WordsEn_proposed_tmthrgd CDF](graphs/shoco_WordsEn_proposed_tmthrgd_cdf.png)

![shoco_WordsEn_proposed_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_WordsEn_proposed_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_FilePath_tmthrgd`

![shoco_FilePath_tmthrgd CDF](graphs/shoco_FilePath_tmthrgd_cdf.png)

![shoco_FilePath_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_FilePath_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_Emails_tmthrgd`

![shoco_Emails_tmthrgd CDF](graphs/shoco_Emails_tmthrgd_cdf.png)

![shoco_Emails_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_Emails_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_FilePath_proposed_tmthrgd`

![shoco_FilePath_proposed_tmthrgd CDF](graphs/shoco_FilePath_proposed_tmthrgd_cdf.png)

![shoco_FilePath_proposed_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_FilePath_proposed_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

### `shoco_Emails_proposed_tmthrgd`

![shoco_Emails_proposed_tmthrgd CDF](graphs/shoco_Emails_proposed_tmthrgd_cdf.png)

![shoco_Emails_proposed_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_Emails_proposed_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

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

