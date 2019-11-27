<?php
/**
 * @author 刘国君
 * @version 1.0
 */

//$client=stream_socket_client('tcp://106.75.154.221:8350');
$client=stream_socket_client('tcp://127.0.0.1:10034');
fwrite($client,'{G01:869975034791700:A0:B1:C00001:D10000:3700:19:30:01:00000B}');
echo fread($client,30);

//fwrite($client,'{A:30:01:P1000H4000L1000}');
//echo fread($client,5);

//fwrite($client,'{A:"111.222.333.444","1234"}');
//echo fread($client,1024);