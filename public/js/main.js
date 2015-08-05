$(document).foundation();
var socket = io.connect();

console.log("Lol");

socket.on('connect', function(){
  console.log("Connected!");
  var html = Handlebars.templates.demo();
  $("#alert").html(html);
  $(document).foundation('alert', 'reflow');
});

$(window).on('beforeunload', function(){
  socket.close();
});
