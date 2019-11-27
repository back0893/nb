<?php
/**
 * @author 刘国君
 * @version 1.0
 */

//$client=stream_socket_client('tcp://106.75.154.221:8350');
$client=stream_socket_client('tcp://127.0.0.1:10034');
//fwrite($client,'{G01:869975031018818:A0:B1:C00001:D10000:3700:19:30:01:00000B}');
//echo fread($client,30);

//fwrite($client,'{A:30:01:P1000H4000L1000}');
//echo fread($client,5);
$hex='5B 00 00 00 92 00 00 06 82 94 00 01 33 EF B8 01 00 00 00 00 00 27 0F D4 C1 41 31 32 33 34 35 00 00 00 00 00 00 00 00 00 00 00 00 00 02 94 01 00 00 00 5C 01 00 02 00 00 00 00 5A 01 AC 3F 40 12 3F FA A1 00 00 00 00 5A 01 AC 4D 50 03 73 6D 61 6C 6C 63 68 69 00 00 00 00 00 00 00 00 31 32 33 34 35 36 37 38 39 30 31 00 00 00 00 00 00 00 00 00 31 32 33 34 35 36 40 71 71 2E 63 6F 6D 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 BA D8 5D';
fwrite($client,hex2bin(str_replace(' ','',$hex)));
echo fread($client,1024);