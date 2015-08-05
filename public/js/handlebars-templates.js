(function() {
  var template = Handlebars.template, templates = Handlebars.templates = Handlebars.templates || {};
templates['demo'] = template({"compiler":[6,">= 2.0.0-beta.1"],"main":function(depth0,helpers,partials,data) {
    return "<div data-alert class=\"alert-box success radius\">\n  Successfully connected to server.\n  <a href=\"#\" class=\"close\">&times;</a>\n</div>\n";
},"useData":true});
})();