<?php
/**
 * @author 刘国君
 * @version 1.0
 */

$client=stream_socket_client('tcp://127.0.0.1:10034');
fwrite($client,'hi,wrold!'."\n");
echo fread($client,9);