<!DOCTYPE html>
	<html>
	<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<meta name="viewport" content="width=device-width,initial-scale=1,user-scalable=no" />
	<style>
	.item_container {
padding:16px 16px 30px;
	font-family:"Helvetica Neue",Helvetica,"lucida Grande",Verdana,"Microsoft YaHei";
	font-size:14px;
	line-height:1.4;
	border-bottom:1px solid #eaeaea;
	}
.item_logoLink {
float:left;
      margin-right:10px;
      margin-bottom:10px;
      font-size:0;
      line-height:0;
}
.item_logo {
border:1px solid #c8c8c8;
}
.item_header {
	margin-bottom:10px;
}
.item_header_title {
	padding-right:4px;
	font-weight:bold;
color:#df0202;
}
.item_btn {
display:inline-block;
padding:6px 16px;
color:#fff;
      background-color:#339dec;
      border-radius:2px;
      text-decoration:none!important;
      font-size:14px;
      *zoom:1;
}
.item_btn:hover {
	text-decoration:none;
}
.item_footer {
	margin-top:16px;
}
.item_header,
	.item_content {
overflow:hidden;
	 *zoom:1;
	}
.item_content_main,
	.item_content_tips {
		margin-bottom:6px;
	}
.item_header_datetime,
	.item_content_tips {
color:#aeaeae;
      font-size:12px;
	}
@media (max-width:600px) {
	.item_container {
		font-size:17px;
	}
	.item_header_title {
display:block;
color:#000;
      font-weight:normal;
	}
	.item_header_datetime {
display:block;
	margin-top:3px;
	}
	.item_content {
overflow:visible;
clear:both;
	}
	.item_footer {
display:none!important;
	}
	.item_btn {
padding:8px 20px;
	font-size:15px;
	}
	.item_header_datetime,
		.item_content_tips {
			font-size:15px;
		}
}
</style>
<style>
.clearfix:after {
	content: ".";
	display: block;
	height: 0;
	clear: both;
	visibility: hidden;
}
.clearfix {
	*display:inline-block;
	*zoom:100%;	
}
.open_email{
	background:url(http://imgcache.qq.com/bossweb/pay/images/mailmsg/email_bg.png) no-repeat 0 -35px;
	width:760px;
	padding:10px;
	font-family: Tahoma,"宋体";
	margin:0 auto;
	margin-bottom:20px;	
	text-align:left;
}
.open_email a:link,.open_email a:visited {
    color:#295394;
    text-decoration:none !important;
}
.open_email a:active,.open_email a:hover {
    color:#000 ;
    text-decoration:underline !important;
}

a.lv:link, a.lv:visited {
	color:#1969e2;
	text-decoration:underline;
}
a.lv:active,a.lv:hover {
	color:#000;
	text-decoration:underline;
}
.open_email .img_left{
	float:left;
	width:125px;
	text-align:left;
}
.open_email .word_right{
	float:right;
	width:615px;
}
.open_email h5,.open_email h6{
	font-size:14px;
	margin:0px;
	padding-top:2px;
	line-height:21px;
}
.open_email h5{
	color:#df0202;
	padding-bottom:10px;
}
.open_email h6{
	padding-bottom:2px;
}
.open_email h5 span,.open_email p{
	font-size:12px;
	color:#808080;
	font-weight:normal;
	margin:0;
	padding:0;
	line-height:21px;
}
.email_btn{
	background:url(http://imgcache.qq.com/bossweb/pay/images/mailmsg/email_bg.png) no-repeat 0 0;
	width:84px;
	height:29px;
	line-height:29px;
	border:0;
	padding:0px;
	cursor:pointer;
	color:#023a65;
	font-size:14px;
	vertical-align:middle;
	text-align:center;
	margin:10px 0 6px;
	display:block;
text-decoration:none;
}
.email_hr{
	background:url(http://imgcache.qq.com/bossweb/pay/images/mailmsg/email_bg.png) no-repeat -100px 0px ;
	height:4px;
	margin:8px 0;
	>overflow:hidden;
}
</style>
</head>
<body>
<div>
{{.Body}}	
</div>
</body>
</html>