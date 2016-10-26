$(function(){
  var newHash      = "";


  $("nav").delegate("a", "click", function() {
      window.location.hash = $(this).attr("href");
      return false;
  });

  $(window).bind('hashchange', function(){
      newHash = window.location.hash.substring(1)?window.location.hash.substring(1):"home";
      if (newHash) {
        $("div.container > div.main-content").hide();
        $("div#"+newHash).show()
        $("nav li").removeClass("active");
        $("nav li a[href=#"+newHash+"]").parent().addClass("active");
      };
  });
  $(window).trigger('hashchange');

   $('#search').on('click', function(e){
       e.preventDefault(); // prevent the default click action
       var q = $('#query').val();
        if( !q.match('^#') &&  !q.match('^@')) {
          alert('the input should start with # or @');
          return;
        }
        console.log("query string "+ q)
       $.ajax({
           url: '/twitter/search?h='+(q.match('^#')?"true":"false")+'&q='+q.substring(1),
           success: function (response) {
               console.log('https://twitter.com/'+response.user.screen_name+'/status/'+response.id_str);
               var tweet = document.getElementById("tweet");
               var id =response.id_str;
               tweet.setAttribute("tweetID",id);
               tweet.innerHTML="";
               twttr.widgets.createTweet(
                id, tweet,
                {
                  conversation : 'none',    // or all
                  cards        : 'hidden',  // or visible
                  linkColor    : '#cc0000', // default is blue
                  theme        : 'light'    // or dark
                })
              .then (function (el) {
                $body.removeClass("loading");
              });
           },
           error: function (response) {
               alert("An error occured. Search someting else. ")
               $body.removeClass("loading");
           },
       });
   });
   $body = $("body");
   $(document).on({
     ajaxStart: function() { $body.addClass("loading");    }
   });
});
