/* jshint esversion: 6 */

function decorateTags(selector) {
	$(selector).each(function () {
		let el = $(this);
		let tags = el.text().split(/\r?\n/)
			.map(w => w.trim()).filter(Boolean);
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
	decorateTags('td[data-heading*="Tags"] div, p[data-heading*="Tags"]');
});