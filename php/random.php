<?php 


/**
 * 获取随机字符串 支持获取指定随机字符串的类型
 * $num => 返回随机字符串的个数
 * $type => 返回随机字符串的字符类型
 */

function getRandom($num = 6,$type = 'all'){

    $str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    switch($type){
        case 'str':
            $str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
            break;
        case 'num':
            $str = '01234567';
            break;
        case 'low':
            $str = "abcdefghijklmnopqrstuvwxyz";
            break;
        case 'upper':
            $str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
            break;
    }

    $res = '';
    for($i=0;$i<$num;$i++){
        $index = rand(0,strlen($str));
        $res .= $str[$index];
    }


    return $res;

}



