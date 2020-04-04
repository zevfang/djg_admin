;
layui.define(function(e) {
	var a = layui.admin;
	layui.use(["admin", "carousel"],
	function() {
		var e = layui.$,
		a = (layui.admin, layui.carousel),
		l = layui.element,
		t = layui.device();
		e(".layadmin-carousel").each(function() {
			var l = e(this);
			a.render({
				elem: this,
				width: "100%",
				arrow: "none",
				interval: l.data("interval"),
				autoplay: l.data("autoplay") === !0,
				trigger: t.ios || t.android ? "click": "hover",
				anim: l.data("anim")
			})
		}),
		l.render("progress")
	}),
	layui.use(["carousel", "echarts"],
	function() {
		var e = layui.$,
		a = (layui.carousel, layui.echarts),
		l = [],
		t = [{
			title: {
				subtext: "完全实况球员数据",
				textStyle: {
					fontSize: 14
				}
			},
			tooltip: {
				trigger: "axis"
			},
			legend: {
				x: "left",
				data: ["罗纳尔多", "舍普琴科"]
			},
			polar: [{
				indicator: [{
					text: "进攻",
					max: 100
				},
				{
					text: "防守",
					max: 100
				},
				{
					text: "体能",
					max: 100
				},
				{
					text: "速度",
					max: 100
				},
				{
					text: "力量",
					max: 100
				},
				{
					text: "技巧",
					max: 100
				}],
				radius: 130
			}],
			series: [{
				type: "radar",
				center: ["50%", "50%"],
				itemStyle: {
					normal: {
						areaStyle: {
							type: "default"
						}
					}
				},
				data: [{
					value: [97, 42, 88, 94, 90, 86],
					name: "舍普琴科"
				},
				{
					value: [97, 32, 74, 95, 88, 92],
					name: "罗纳尔多"
				}]
			}]
		}],
		i = e("#LAY-index-pageone").children("div"),
		n = function(e) {
			l[e] = a.init(i[e], layui.echartsTheme),
			l[e].setOption(t[e]),
			window.onresize = l[e].resize
		};
		i[0] && n(0)
	}),
	layui.use(["carousel", "echarts"],
	function() {
		var e = layui.$,
		a = (layui.carousel, layui.echarts),
		l = [],
		t = [{
			tooltip: {
				trigger: "axis"
			},
			calculable: !0,
			legend: {
				data: ["访问量", "下载量", "平均访问量"]
			},
			xAxis: [{
				type: "category",
				data: ["1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"]
			}],
			yAxis: [{
				type: "value",
				name: "访问量",
				axisLabel: {
					formatter: "{value} 万"
				}
			},
			{
				type: "value",
				name: "下载量",
				axisLabel: {
					formatter: "{value} 万"
				}
			}],
			series: [{
				name: "访问量",
				type: "line",
				data: [900, 850, 950, 1e3, 1100, 1050, 1e3, 1150, 1250, 1370, 1250, 1100]
			},
			{
				name: "下载量",
				type: "line",
				yAxisIndex: 1,
				data: [850, 850, 800, 950, 1e3, 950, 950, 1150, 1100, 1240, 1e3, 950]
			},
			{
				name: "平均访问量",
				type: "line",
				data: [870, 850, 850, 950, 1050, 1e3, 980, 1150, 1e3, 1300, 1150, 1e3]
			}]
		}],
		i = e("#LAY-index-pagetwo").children("div"),
		n = function(e) {
			l[e] = a.init(i[e], layui.echartsTheme),
			l[e].setOption(t[e]),
			window.onresize = l[e].resize
		};
		i[0] && n(0)
	}),
	a.events.replyNote = function(e) {
		var a = e.data("id");
		layer.prompt({
			title: "回复留言 ID:" + a,
			formType: 2
		},
		function(e, a) {
			layer.msg("得到：" + e),
			layer.close(a)
		})
	},
	e("sample", {})
});