<body>
	<div class="item_container">
		<a class="item_logoLink" href="http://wap.douban.com/people/44407326/about" target="_blank">
			<img class="item_logo" width="93" height="68" src="http://img3.douban.com/icon/u44407326-24.jpg" />
		</a>
		<div class="item_header">
			<span class="item_header_title">
                {{.Subject}}</span>
			<span class="item_header_datetime">({{.Datetime}})</span>
		</div>
		<div class="item_content">
			<div class="item_content_main">
        			{{.Content}}
			</div>
			<div class="item_content_tips">
                点击按钮查看详细信息</div>
			<div class="item_content_action">
				<a class="item_btn" href="{{.Url}}" target="_blank">进入查看</a>
			</div>
		</div>
		<div class="item_footer">
			
		</div>
	</div>
</body>