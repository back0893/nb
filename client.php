<?php
/**
 * @author 刘国君
 * @version 1.0
 */

$client=stream_socket_client('tcp://127.0.0.1:10034');
fwrite($client,'{G01:869975031018818:A0:B1:C0.001V:D10.01V:3.7V:30:01:000034}!');
echo fread($client,30);