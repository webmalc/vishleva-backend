function decorateTags() {
	$('td[data-heading*="Tags"] div').each(function () {
		let el = $(this);
		let tags = el.text().replace(/[\t\n]+/g, '').replace(/ +(?= )/g, '')
			.trim().split(" ");
		console.log(tags);
		let content = "";
		tags.forEach(function (tag) {
			if (tag) {
				content += '<span class="v-tag">' + tag + '</span>';
			}
		});
		el.html(content);
	});
}

$("document").ready(function () {
	// bool
	$(".qor-table__content:contains('true')").css("color", "green");
	$(".qor-table__content:contains('false')").css("color", "maroon");

	// tags
	decorateTags();
});