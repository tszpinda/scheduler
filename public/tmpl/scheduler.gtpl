<html>
<head>
   <head>
      <title>Search</title>
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <link rel="shortcut icon" href="/ui/static/favicon.ico">
      
      <link href="/ui/static/bootstrap/css/main.css" rel="stylesheet" media="screen">
      <link href="/ui/static/scheduler/dhtmlxscheduler.css" rel="stylesheet" media="screen">
      <link href="/ui/static/dataview/dhtmlxdataview.css" rel="stylesheet" media="screen">

      <script src="/ui/static/js/jquery.js"></script>
      <script src="/ui/static/bootstrap/js/bootstrap.js"></script>
      <script src="/ui/static/js/main.js"></script>
      <script src="/ui/static/scheduler/dhtmlxscheduler.js"></script>
      <script src="/ui/static/scheduler/dhtmlxscheduler-units.js"></script>
      <script src='/ui/static/scheduler/dhtmlxscheduler_outerdrag.js' type="text/javascript"></script>

      <script src="/ui/static/dataview/dhtmlxdataview.js"></script>

      <script type="text/javascript" charset="utf-8">

         $(function(){
            var xhtmlDateFormat = "%Y-%m-%d %H:%i";
            dhtmlx.compat("dnd");

            var sections = scheduler.serverList("type");
                 
            scheduler.createUnitsView({
                name:"unit",
                property:"type",
                list:sections});
      
            //section config
            scheduler.config.first_hour = 8;
            scheduler.locale.labels.unit_tab = "Trucks"


            $("#scheduler").dhx_scheduler({
                 //xml_date:"%d-%m-%Y %H:%i",
               xml_date:xhtmlDateFormat,
               date:new Date(),
                 mode:"unit",
                 details_on_create:true,
                 details_on_dblclick:true
            });
          
             scheduler.load("/api/events/1", "json");

/*
             scheduler.renderEvent = function(container, ev) {
                var container_width = container.style.width; // e.g. "105px"
             
                // move section
                var html = "<div class='dhx_event_move my_event_move' style='width: " + 
                container_width + "'></div>";
             
                // container for event contents
                html+= "<div class='my_event_body'>";
                html += "<span class='event_date'>";
                //two options here:show only start date for short events or start+end for long
                if ((ev.end_date - ev.start_date)/60000>40){//if event is longer than 40 minutes
                    html += scheduler.templates.event_header(ev.start_date, ev.end_date, ev);
                    html += "</span><br/>";
                } else {
                    html += scheduler.templates.event_date(ev.start_date) + "</span>";
                }
                // displaying event text
                html += "<span>" + scheduler.templates.event_text(ev.start_date,ev.end_date,ev)+
                "</span>" + "</div>";
             
                // resize section
                html += "<div class='dhx_event_resize my_event_resize' style='width: " +
                container_width + "'></div>";
             
                container.innerHTML = html;
                return true; //required, true - display a custom form, false - the default form
            };*/

            var data = new dhtmlXDataView({
              container:"holdingArea",
              type:{
                template:"#name#",
                height:35,
                template_loading:"Loading...",                
              },
              drag: true
            });
            data.load("/ui/static/holdingArea.json", "json");

            scheduler.attachEvent("onExternalDragIn", function(id, source, mouseEvent) {
              var context = dhtmlx.DragControl.getContext();
              var newDelivery = context.from.get(context.source[0]);
              //scheduler.getEvent(id).text = label;

              return true;
            });

         });
      </script>
   </head>
   <body>
            <div id="holdingArea" ></div>
            <div id="scheduler" class="dhx_cal_container">
               <div class="dhx_cal_navline">
                  <div class="dhx_cal_prev_button">&nbsp;</div>
                  <div class="dhx_cal_next_button">&nbsp;</div>
                  <div class="dhx_cal_today_button"></div>
                  <div class="dhx_cal_date"></div>
                  <div class="dhx_cal_tab" name="day_tab" style="right:204px;"></div>
                  <div class="dhx_cal_tab" name="week_tab" style="right:140px;"></div>
                  <div class="dhx_cal_tab" name="month_tab" style="right:76px;"></div>
                  <div class="dhx_cal_tab" name="unit_tab" style="right:280px;"></div>
               </div>
               <div class="dhx_cal_header">
               </div>
               <div class="dhx_cal_data">
               </div>
            </div>
   </body>
</html>
