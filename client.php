<?php
/**
 * @author 刘国君
 * @version 1.0
 */

$client=stream_socket_client('tcp://120.132.93.231:8084');
//$client=stream_socket_client('tcp://106.75.154.221:8350');
//$client=stream_socket_client('tcp://127.0.0.1:10034');
while(1){
    fwrite($client,'{G01:869975034791700:A1:B1:C01075:D44816:2419:27:30:01:201138}');
    echo fread($client,30);

    sleep(5);
}

//fwrite($client,'{A:30:01:P1000H4000L1000}');
//echo fread($client,5);

//fwrite($client,'{A:"111.222.333.444","1234"}');
//echo fread($client,1024);