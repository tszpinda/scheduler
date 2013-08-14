<html>
<head>
   <head>
      <title>Search</title>
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <link rel="shortcut icon" href="/ui/static/favicon.ico">
      <link href="/ui/static/bootstrap/css/bootstrap.css" rel="stylesheet" media="screen">
      <link href="/ui/static/bootstrap/css/bootstrap-responsive.css" rel="stylesheet" media="screen">
      <link href="/ui/static/bootstrap/css/main.css" rel="stylesheet" media="screen">
      <link href="/ui/static/scheduler/dhtmlxscheduler.css" rel="stylesheet" media="screen">
      <script src="/ui/static/js/jquery.js"></script>
      <script src="/ui/static/bootstrap/js/bootstrap.js"></script>
      <script src="/ui/static/js/main.js"></script>

      <script src="/ui/static/scheduler/dhtmlxscheduler.js"></script>
      <script src="/ui/static/scheduler/dhtmlxscheduler-units.js"></script>
      <script type="text/javascript" charset="utf-8">

         $(function(){
                        //scheduler.config.xml_date="%d-%m-%Y %H:%i";
            //scheduler.init('scheduler',new Date(),"week");
            //scheduler.load("/scheduler/event/1", "json");            

             $("#scheduler").dhx_scheduler({
                 xml_date:"%d-%m-%Y %H:%i",
                 date:new Date(),
                 mode:"week",
                 details_on_create:true,
                 details_on_dblclick:true
             });
          
             scheduler.load("/scheduler/event/1", "json");
             
         });
         //function init() {
            
            

      /*      var sections=[333
               {key:1, label:"James Smith"},
               {key:2, label:"John Williams"},
               {key:3, label:"David Miller"},
               {key:4, label:"Linda Brown"}
            ];
            
            scheduler.locale.labels.unit_tab = "Trucks"
            scheduler.locale.labels.section_custom="Assigned to";
            scheduler.config.first_hour = 8;
            scheduler.config.multi_day = true;
            scheduler.config.details_on_create=true;
            scheduler.config.details_on_dblclick=true;
            scheduler.config.xml_date="%Y-%m-%d %H:%i";
            scheduler.templates.event_class=function(s,e,ev){ return ev.custom?"custom":""; };
            scheduler.config.lightbox.sections=[   
               {name:"description", height:130, map_to:"text", type:"textarea" , focus:true},
               {name:"custom", height:23, type:"select", options:sections, map_to:"section_id" },
               {name:"time", height:72, type:"time", map_to:"auto"}
            ]
      
            scheduler.createUnitsView("unit","section_id",sections);
      
*/
            
         //}
      </script>
   </head>
   <body>
      <b>Scheduler</b>
      <div id="scheduler" class="dhx_cal_container" style='width:100%; height:100%;'>
         <div class="dhx_cal_navline">
            <div class="dhx_cal_prev_button">&nbsp;</div>
            <div class="dhx_cal_next_button">&nbsp;</div>
            <div class="dhx_cal_today_button"></div>
            <div class="dhx_cal_date"></div>
            <div class="dhx_cal_tab" name="day_tab" style="right:204px;"></div>
            <div class="dhx_cal_tab" name="week_tab" style="right:140px;"></div>
            <div class="dhx_cal_tab" name="month_tab" style="right:76px;"></div>
         </div>
         <div class="dhx_cal_header">
         </div>
         <div class="dhx_cal_data">
         </div>
      </div>
   </body>
</html>
