type = ['', 'info', 'success', 'warning', 'danger'];


demo = {

  initCirclePercentage: function () {

    $('#chartDashboard, #chartOrders, #chartNewVisitors, #chartSubscriptions, #chartDashboardDoc, #chartOrdersDoc').easyPieChart({
      lineWidth: 6,
      size: 160,
      scaleColor: false,
      trackColor: 'rgba(255,255,255,.25)',
      barColor: '#FFFFFF',
      animate: ({ duration: 5000, enabled: true })

    });


  },

  initGoogleMaps: function () {

    // Satellite Map
    var myLatlng = new google.maps.LatLng(40.748817, -73.985428);
    var mapOptions = {
      zoom: 3,
      scrollwheel: false, //we disable de scroll over the map, it is a really annoing when you scroll through page
      center: myLatlng,
      mapTypeId: google.maps.MapTypeId.SATELLITE
    }

    var map = new google.maps.Map(document.getElementById("satelliteMap"), mapOptions);

    var marker = new google.maps.Marker({
      position: myLatlng,
      title: "Satellite Map!"
    });

    marker.setMap(map);


  },

  initSmallGoogleMaps: function () {

    // Regular Map
    var myLatlng = new google.maps.LatLng(40.748817, -73.985428);
    var mapOptions = {
      zoom: 8,
      center: myLatlng,
      scrollwheel: false, //we disable de scroll over the map, it is a really annoing when you scroll through page
    }

    var map = new google.maps.Map(document.getElementById("regularMap"), mapOptions);

    var marker = new google.maps.Marker({
      position: myLatlng,
      title: "Regular Map!"
    });

    marker.setMap(map);


    // Custom Skin & Settings Map
    var myLatlng = new google.maps.LatLng(40.748817, -73.985428);
    var mapOptions = {
      zoom: 13,
      center: myLatlng,
      scrollwheel: false, //we disable de scroll over the map, it is a really annoing when you scroll through page
      disableDefaultUI: true, // a way to quickly hide all controls
      zoomControl: true,
      styles: [{ "featureType": "water", "stylers": [{ "saturation": 43 }, { "lightness": -11 }, { "hue": "#0088ff" }] }, { "featureType": "road", "elementType": "geometry.fill", "stylers": [{ "hue": "#ff0000" }, { "saturation": -100 }, { "lightness": 99 }] }, { "featureType": "road", "elementType": "geometry.stroke", "stylers": [{ "color": "#808080" }, { "lightness": 54 }] }, { "featureType": "landscape.man_made", "elementType": "geometry.fill", "stylers": [{ "color": "#ece2d9" }] }, { "featureType": "poi.park", "elementType": "geometry.fill", "stylers": [{ "color": "#ccdca1" }] }, { "featureType": "road", "elementType": "labels.text.fill", "stylers": [{ "color": "#767676" }] }, { "featureType": "road", "elementType": "labels.text.stroke", "stylers": [{ "color": "#ffffff" }] }, { "featureType": "poi", "stylers": [{ "visibility": "off" }] }, { "featureType": "landscape.natural", "elementType": "geometry.fill", "stylers": [{ "visibility": "on" }, { "color": "#b8cb93" }] }, { "featureType": "poi.park", "stylers": [{ "visibility": "on" }] }, { "featureType": "poi.sports_complex", "stylers": [{ "visibility": "on" }] }, { "featureType": "poi.medical", "stylers": [{ "visibility": "on" }] }, { "featureType": "poi.business", "stylers": [{ "visibility": "simplified" }] }]

    }

    var map = new google.maps.Map(document.getElementById("customSkinMap"), mapOptions);

    var marker = new google.maps.Marker({
      position: myLatlng,
      title: "Custom Skin & Settings Map!"
    });

    marker.setMap(map);

  },


  initVectorMap: function () {
    var mapData = {
      "AU": 760,
      "BR": 550,
      "CA": 120,
      "DE": 1300,
      "FR": 540,
      "GB": 690,
      "GE": 200,
      "IN": 200,
      "RO": 600,
      "RU": 300,
      "US": 2920,
    };

    $('#worldMap').vectorMap({
      map: 'world_mill_en',
      backgroundColor: "transparent",
      zoomOnScroll: false,
      regionStyle: {
        initial: {
          fill: '#e4e4e4',
          "fill-opacity": 0.9,
          stroke: 'none',
          "stroke-width": 0,
          "stroke-opacity": 0
        }
      },

      series: {
        regions: [{
          values: mapData,
          scale: ["#AAAAAA", "#444444"],
          normalizeFunction: 'polynomial'
        }]
      },
    });
  },

  initFullScreenGoogleMap: function () {
    var myLatlng = new google.maps.LatLng(40.748817, -73.985428);
    var mapOptions = {
      zoom: 13,
      center: myLatlng,
      scrollwheel: false, //we disable de scroll over the map, it is a really annoing when you scroll through page
      styles: [{ "featureType": "water", "stylers": [{ "saturation": 43 }, { "lightness": -11 }, { "hue": "#0088ff" }] }, { "featureType": "road", "elementType": "geometry.fill", "stylers": [{ "hue": "#ff0000" }, { "saturation": -100 }, { "lightness": 99 }] }, { "featureType": "road", "elementType": "geometry.stroke", "stylers": [{ "color": "#808080" }, { "lightness": 54 }] }, { "featureType": "landscape.man_made", "elementType": "geometry.fill", "stylers": [{ "color": "#ece2d9" }] }, { "featureType": "poi.park", "elementType": "geometry.fill", "stylers": [{ "color": "#ccdca1" }] }, { "featureType": "road", "elementType": "labels.text.fill", "stylers": [{ "color": "#767676" }] }, { "featureType": "road", "elementType": "labels.text.stroke", "stylers": [{ "color": "#ffffff" }] }, { "featureType": "poi", "stylers": [{ "visibility": "off" }] }, { "featureType": "landscape.natural", "elementType": "geometry.fill", "stylers": [{ "visibility": "on" }, { "color": "#b8cb93" }] }, { "featureType": "poi.park", "stylers": [{ "visibility": "on" }] }, { "featureType": "poi.sports_complex", "stylers": [{ "visibility": "on" }] }, { "featureType": "poi.medical", "stylers": [{ "visibility": "on" }] }, { "featureType": "poi.business", "stylers": [{ "visibility": "simplified" }] }]

    }
    var map = new google.maps.Map(document.getElementById("map"), mapOptions);

    var marker = new google.maps.Marker({
      position: myLatlng,
      title: "Hello World!"
    });

    // To add the marker to the map, call setMap();
    marker.setMap(map);
  },

  initOverviewDashboardDoc: function () {
    /*  **************** Chart Total Earnings - single line ******************** */

    var dataPrice = {
      labels: ['Jan', 'Feb', 'Mar', 'April', 'May', 'June'],
      series: [
        [230, 340, 400, 300, 570, 500, 800]
      ]
    };

    var optionsPrice = {
      showPoint: false,
      lineSmooth: true,
      height: "210px",
      axisX: {
        showGrid: false,
        showLabel: true
      },
      axisY: {
        offset: 40,
        showGrid: false
      },
      low: 0,
      high: 'auto',
      classNames: {
        line: 'ct-line ct-green'
      }
    };

    Chartist.Line('#chartTotalEarningsDoc', dataPrice, optionsPrice);

    /*  **************** Chart Subscriptions - single line ******************** */

    var dataDays = {
      labels: ['M', 'T', 'W', 'T', 'F', 'S', 'S'],
      series: [
        [60, 50, 30, 50, 70, 60, 90, 100]
      ]
    };

    var optionsDays = {
      showPoint: false,
      lineSmooth: true,
      height: "210px",
      axisX: {
        showGrid: false,
        showLabel: true
      },
      axisY: {
        offset: 40,
        showGrid: false
      },
      low: 0,
      high: 'auto',
      classNames: {
        line: 'ct-line ct-red'
      }
    };

    Chartist.Line('#chartTotalSubscriptionsDoc', dataDays, optionsDays);
  },

  initOverviewDashboard: function () {

    /*  **************** Chart Total Earnings - single line ******************** */

    var dataPrice = {
      labels: ['Jan', 'Feb', 'Mar', 'April', 'May', 'June'],
      series: [
        [230, 340, 400, 300, 570, 500, 800]
      ]
    };

    var optionsPrice = {
      showPoint: false,
      lineSmooth: true,
      height: "210px",
      axisX: {
        showGrid: false,
        showLabel: true
      },
      axisY: {
        offset: 40,
        showGrid: false
      },
      low: 0,
      high: 'auto',
      classNames: {
        line: 'ct-line ct-green'
      }
    };

    Chartist.Line('#chartTotalEarnings', dataPrice, optionsPrice);

    /*  **************** Chart Subscriptions - single line ******************** */

    var dataDays = {
      labels: ['M', 'T', 'W', 'T', 'F', 'S', 'S'],
      series: [
        [60, 50, 30, 50, 70, 60, 90, 100]
      ]
    };

    var optionsDays = {
      showPoint: false,
      lineSmooth: true,
      height: "210px",
      axisX: {
        showGrid: false,
        showLabel: true
      },
      axisY: {
        offset: 40,
        showGrid: false
      },
      low: 0,
      high: 'auto',
      classNames: {
        line: 'ct-line ct-red'
      }
    };

    Chartist.Line('#chartTotalSubscriptions', dataDays, optionsDays);

    /*  **************** Chart Total Downloads - single line ******************** */

    var dataDownloads = {
      labels: ['2009', '2010', '2011', '2012', '2013', '2014'],
      series: [
        [1200, 1000, 3490, 8345, 3256, 2566]
      ]
    };

    var optionsDownloads = {
      showPoint: false,
      lineSmooth: true,
      height: "210px",
      axisX: {
        showGrid: false,
        showLabel: true
      },
      axisY: {
        offset: 40,
        showGrid: false
      },
      low: 0,
      high: 'auto',
      classNames: {
        line: 'ct-line ct-orange'
      }
    };

    Chartist.Line('#chartTotalDownloads', dataDownloads, optionsDownloads);
  },

  initStatsDashboard: function () {

    var dataSales = {
      labels: ['9:00AM', '12:00AM', '3:00PM', '6:00PM', '9:00PM', '12:00PM', '3:00AM', '6:00AM'],
      series: [
        [287, 385, 490, 562, 594, 626, 698, 895, 952],
        [67, 152, 193, 240, 387, 435, 535, 642, 744],
        [23, 113, 67, 108, 190, 239, 307, 410, 410]
      ]
    };

    var optionsSales = {
      lineSmooth: false,
      low: 0,
      high: 1000,
      showArea: true,
      height: "245px",
      axisX: {
        showGrid: false,
      },
      lineSmooth: Chartist.Interpolation.simple({
        divisor: 3
      }),
      showLine: true,
      showPoint: false,
    };

    var responsiveSales = [
      ['screen and (max-width: 640px)', {
        axisX: {
          labelInterpolationFnc: function (value) {
            return value[0];
          }
        }
      }]
    ];

    Chartist.Line('#chartHours', dataSales, optionsSales, responsiveSales);


    var data = {
      labels: ['Jan', 'Feb', 'Mar', 'Apr', 'Mai', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
      series: [
        [542, 543, 520, 680, 653, 753, 326, 434, 568, 610, 756, 895],
        [230, 293, 380, 480, 503, 553, 600, 664, 698, 710, 736, 795]
      ]
    };

    var options = {
      seriesBarDistance: 10,
      axisX: {
        showGrid: false
      },
      height: "245px"
    };

    var responsiveOptions = [
      ['screen and (max-width: 640px)', {
        seriesBarDistance: 5,
        axisX: {
          labelInterpolationFnc: function (value) {
            return value[0];
          }
        }
      }]
    ];

    Chartist.Line('#chartActivity', data, options, responsiveOptions);

    Chartist.Pie('#chartPreferences', {
      labels: ['62%', '32%', '6%'],
      series: [62, 32, 6]
    });
  },

  initChartsPage: function () {
    /*  **************** 24 Hours Performance - single line ******************** */

    var dataPerformance = {
      labels: ['6pm', '9pm', '11pm', '2am', '4am', '8am', '2pm', '5pm', '8pm', '11pm', '4am'],
      series: [
        [1, 6, 8, 7, 4, 7, 8, 12, 16, 17, 14, 13]
      ]
    };

    var optionsPerformance = {
      showPoint: false,
      lineSmooth: true,
      height: "200px",
      axisX: {
        showGrid: false,
        showLabel: true
      },
      axisY: {
        offset: 40,

      },
      low: 0,
      high: 16,
      height: "250px"
    };


    Chartist.Line('#chartPerformance', dataPerformance, optionsPerformance);

    /*   **************** 2014 Sales - Bar Chart ********************    */

    var data = {
      labels: ['Jan', 'Feb', 'Mar', 'Apr', 'Mai', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
      series: [
        [542, 443, 320, 780, 553, 453, 326, 434, 568, 610, 756, 895],
        [412, 243, 280, 580, 453, 353, 300, 364, 368, 410, 636, 695]
      ]
    };

    var options = {
      seriesBarDistance: 10,
      axisX: {
        showGrid: false
      },
      height: "250px"
    };

    var responsiveOptions = [
      ['screen and (max-width: 640px)', {
        seriesBarDistance: 5,
        axisX: {
          labelInterpolationFnc: function (value) {
            return value[0];
          }
        }
      }]
    ];

    Chartist.Bar('#chartActivity', data, options, responsiveOptions);


    /*  **************** NASDAQ: AAPL - single line with points ******************** */

    var dataStock = {
      labels: ['\'07', '\'08', '\'09', '\'10', '\'11', '\'12', '\'13', '\'14', '\'15'],
      series: [
        [22.20, 34.90, 42.28, 51.93, 62.21, 80.23, 62.21, 82.12, 102.50, 107.23], [22.20, 34.90, 42.28]
      ]
    };

    var optionsStock = {
      lineSmooth: false,
      height: "200px",
      axisY: {
        offset: 40,
        labelInterpolationFnc: function (value) {
          return '$' + value;
        }

      },
      low: 10,
      height: "250px",
      high: 110,
      classNames: {
        point: 'ct-point ct-green',
        line: 'ct-line ct-green'
      }
    };

    Chartist.Line('#chartStock', dataStock, optionsStock);

    /*  **************** Views  - barchart ******************** */

    var dataViews = {
      labels: ['Jan', 'Feb', 'Mar', 'Apr', 'Mai', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
      series: [
        [542, 443, 320, 780, 553, 453, 326, 434, 568, 610, 756, 895]
      ]
    };

    var optionsViews = {
      seriesBarDistance: 10,
      classNames: {
        bar: 'ct-bar'
      },
      axisX: {
        showGrid: false,

      },
      height: "250px"

    };

    var responsiveOptionsViews = [
      ['screen and (max-width: 640px)', {
        seriesBarDistance: 5,
        axisX: {
          labelInterpolationFnc: function (value) {
            return value[0];
          }
        }
      }]
    ];

    Chartist.Bar('#chartViews', dataViews, optionsViews, responsiveOptionsViews);


  },

  showSwal: function (type) {
    if (type == 'basic') {
      swal({
        title: "Here's a message!",
        buttonsStyling: false,
        confirmButtonClass: "btn btn-success btn-fill"
      });

    } else if (type == 'update-message') {
      swal({
        title: "Good job!",
        text: "Successfully updated service to the database",
        buttonsStyling: false,
        confirmButtonClass: "btn btn-success btn-fill",
        type: "success"
      });
    } else if (type == 'title-and-text') {
      swal({
        title: "Here's a message!",
        text: "It's pretty, isn't it?",
        buttonsStyling: false,
        confirmButtonClass: "btn btn-info btn-fill"
      });

    } else if (type == 'success-message') {
      swal({
        title: "Good job!",
        text: "Successfully daleted shop from the database",
        buttonsStyling: false,
        confirmButtonClass: "btn btn-success btn-fill",
        type: "success"
      });

    } else if (type == 'warning-message-and-confirmation') {
      swal({
        title: 'Are you sure?',
        text: "You won't be able to revert this!",
        type: 'warning',
        showCancelButton: true,
        confirmButtonClass: 'btn btn-success btn-fill',
        cancelButtonClass: 'btn btn-danger btn-fill',
        confirmButtonText: 'Yes, delete it!',
        buttonsStyling: false
      }).then(function () {
        swal({
          title: 'Deleted!',
          text: 'Your file has been deleted.',
          type: 'success',
          confirmButtonClass: "btn btn-success btn-fill",
          buttonsStyling: false
        })
      });
    } else if (type == 'warning-message-and-cancel') {
      swal({
        title: 'Are you sure?',
        text: 'You will not be able to recover this imaginary file!',
        type: 'warning',
        showCancelButton: true,
        confirmButtonText: 'Yes, delete it!',
        cancelButtonText: 'No, keep it',
        confirmButtonClass: "btn btn-success btn-fill",
        cancelButtonClass: "btn btn-danger btn-fill",
        buttonsStyling: false
      }).then(function () {
        swal({
          title: 'Deleted!',
          text: 'Your imaginary file has been deleted.',
          type: 'success',
          confirmButtonClass: "btn btn-success btn-fill",
          buttonsStyling: false
        })
      }, function (dismiss) {
        // dismiss can be 'overlay', 'cancel', 'close', 'esc', 'timer'
        if (dismiss === 'cancel') {
          swal({
            title: 'Cancelled',
            text: 'Your imaginary file is safe :)',
            type: 'error',
            confirmButtonClass: "btn btn-info btn-fill",
            buttonsStyling: false
          })
        }
      })

    } else if (type == 'custom-html') {
      swal({
        title: 'HTML example',
        buttonsStyling: false,
        confirmButtonClass: "btn btn-success btn-fill",
        html:
          'You can use <b>bold text</b>, ' +
          '<a href="http://github.com">links</a> ' +
          'and other HTML tags'
      });

    } else if (type == 'auto-close') {
      swal({
        title: "Auto close alert!",
        text: "I will close in 2 seconds.",
        timer: 2000,
        showConfirmButton: false
      });
    } else if (type == 'input-field') {
      swal({
        title: 'Input something',
        html: '<div class="form-group">' +
          '<input id="input-field" type="text" class="form-control" />' +
          '</div>',
        showCancelButton: true,
        confirmButtonClass: 'btn btn-success btn-fill',
        cancelButtonClass: 'btn btn-danger btn-fill',
        buttonsStyling: false
      }).then(function (result) {
        swal({
          type: 'success',
          html: 'You entered: <strong>' +
            $('#input-field').val() +
            '</strong>',
          confirmButtonClass: 'btn btn-success btn-fill',
          buttonsStyling: false

        })
      }).catch(swal.noop)
    }
  },

  checkFullPageBackgroundImage: function () {
    $page = $('.full-page');
    image_src = $page.data('image');

    if (image_src !== undefined) {
      image_container = '<div class="full-page-background" style="background-image: url(' + image_src + ') "/>'
      $page.append(image_container);
    }
  },

  initWizard: function () {
    $(document).ready(function () {

      var $validator = $("#wizardForm").validate({
        rules: {
          email: {
            required: true,
            email: true,
            minlength: 5
          },
          first_name: {
            required: false,
            minlength: 5
          },
          last_name: {
            required: false,
            minlength: 5
          },
          website: {
            required: true,
            minlength: 5,
            url: true
          },
          framework: {
            required: false,
            minlength: 4
          },
          cities: {
            required: true
          },
          price: {
            number: true
          }
        }
      });

      // you can also use the nav-pills-[blue | azure | green | orange | red] for a different color of wizard
      $('#wizardCard').bootstrapWizard({
        tabClass: 'nav nav-pills',
        nextSelector: '.btn-next',
        previousSelector: '.btn-back',
        onNext: function (tab, navigation, index) {
          var $valid = $('#wizardForm').valid();

          if (!$valid) {
            $validator.focusInvalid();
            return false;
          }
        },
        onInit: function (tab, navigation, index) {

          //check number of tabs and fill the entire row
          var $total = navigation.find('li').length;
          $width = 100 / $total;

          $display_width = $(document).width();

          if ($display_width < 600 && $total > 3) {
            $width = 50;
          }

          navigation.find('li').css('width', $width + '%');
        },
        onTabClick: function (tab, navigation, index) {
          // Disable the posibility to click on tabs
          return false;
        },
        onTabShow: function (tab, navigation, index) {
          var $total = navigation.find('li').length;
          var $current = index + 1;

          var wizard = navigation.closest('.card-wizard');

          // If it's the last tab then hide the last button and show the finish instead
          if ($current >= $total) {
            $(wizard).find('.btn-next').hide();
            $(wizard).find('.btn-finish').show();
          } else if ($current == 1) {
            $(wizard).find('.btn-back').hide();
          } else {
            $(wizard).find('.btn-back').show();
            $(wizard).find('.btn-next').show();
            $(wizard).find('.btn-finish').hide();
          }
        }
      });
    });

    function onFinishWizard() {
      //here you can do something, sent the form to server via ajax and show a success message with swal

      swal("Good job!", "You clicked the finish button!", "success");
    }
  },

  initFormExtendedSliders: function () {
    // Sliders for demo purpose in refine cards section
    var slider = document.getElementById('sliderRegular');

    noUiSlider.create(slider, {
      start: 40,
      connect: [true, false],
      range: {
        min: 0,
        max: 100
      }
    });

    var slider2 = document.getElementById('sliderDouble');

    noUiSlider.create(slider2, {
      start: [20, 60],
      connect: true,
      range: {
        min: 0,
        max: 100
      }
    });
  },



  initFormExtendedDatetimepickers: function () {

    $('.datetimepicker').datetimepicker({
      format: 'DD/MM/YYYY HH:mm',
      icons: {
        time: "fa fa-clock-o",
        date: "fa fa-calendar",
        up: "fa fa-chevron-up",
        down: "fa fa-chevron-down",
        previous: 'fa fa-chevron-left',
        next: 'fa fa-chevron-right',
        today: 'fa fa-screenshot',
        clear: 'fa fa-trash',
        close: 'fa fa-remove'
      }
    });

    $('.datepicker').datetimepicker({

      format: 'DD/MM/YYYY',    //use this format if you want the 12hours timpiecker with AM/PM toggle
      icons: {
        time: "fa fa-clock-o",
        date: "fa fa-calendar",
        up: "fa fa-chevron-up",
        down: "fa fa-chevron-down",
        previous: 'fa fa-chevron-left',
        next: 'fa fa-chevron-right',
        today: 'fa fa-screenshot',
        clear: 'fa fa-trash',
        close: 'fa fa-remove'
      }
    });
    $('.datepicker1').datetimepicker({
      useCurrent: false,
      format: 'DD/MM/YYYY',    //use this format if you want the 12hours timpiecker with AM/PM toggle
      icons: {
        time: "fa fa-clock-o",
        date: "fa fa-calendar",
        up: "fa fa-chevron-up",
        down: "fa fa-chevron-down",
        previous: 'fa fa-chevron-left',
        next: 'fa fa-chevron-right',
        today: 'fa fa-screenshot',
        clear: 'fa fa-trash',
        close: 'fa fa-remove'
      }
    });
    $(".datepicker1").on("dp.change", function (e) {
      console.log(e.date.format());
      var eventid = $("#eventID").val();
      var eventResourceID = $("#eventresourceID").val();



      var starttimerantevou = $("#startTime").text();
      var starttimerantevoustring = moment(starttimerantevou, 'DD/MM/YYYY HH:mm');
      var momentStringStart = starttimerantevoustring.format('L');

      var endtimerantevou = $("#endTime").text();
      var endtimerantevoustring = moment(endtimerantevou, 'DD/MM/YYYY HH:mm');
      var momentStringEnd = endtimerantevoustring.format('L');



      var result = e.date.diff(starttimerantevoustring, 'days');
      var StartmomentObj = starttimerantevoustring.add(result, 'days').format("YYYY-MM-DD HH:mm");
      var EndmomentObj = endtimerantevoustring.add(result, 'days').format("YYYY-MM-DD HH:mm");



      var formData = {
        'start': StartmomentObj,
        'end': EndmomentObj,
        'id': parseInt(eventid),
        'resource_id': parseInt(eventResourceID),
      };

      $.ajax({
        type: 'post',
        url: '/admin/api/update-calendar-with-hairdresser-id',
        data: JSON.stringify(formData),
        contentType: 'application/json',
        dataType: 'json',

        success: function (events) {

          $('#editfullCalModal').modal('toggle');
          window.location.reload();

        },
        error: function (err) {
          console.log(err);
        }
      });

    });

    $('.timepicker').datetimepicker({
      //          format: 'H:mm',    // use this format if you want the 24hours timepicker
      format: 'H:mm',    //use this format if you want the 12hours timpiecker with AM/PM toggle
      icons: {
        time: "fa fa-clock-o",
        date: "fa fa-calendar",
        up: "fa fa-chevron-up",
        down: "fa fa-chevron-down",
        previous: 'fa fa-chevron-left',
        next: 'fa fa-chevron-right',
        today: 'fa fa-screenshot',
        clear: 'fa fa-trash',
        close: 'fa fa-remove'
      }

    });
  },

  initFullCalendar: function () {
    $calendar = $('#fullCalendar');

    today = new Date();
    y = today.getFullYear();
    m = today.getMonth();
    d = today.getDate();

    $calendar.fullCalendar({
      viewRender: function (view, element) {
        // We make sure that we activate the perfect scrollbar when the view isn't on Month
        if (view.name != 'month') {
          $(element).find('.fc-scroller').perfectScrollbar();
        }
      },
      header: {
        left: 'title',
        center: 'month,agendaWeek,agendaDay',
        right: 'prev,next,today'
      },
      defaultDate: today,
      selectable: true,
      selectHelper: true,
      views: {
        month: { // name of view
          titleFormat: 'MMMM YYYY'
          // other view-specific options here
        },
        week: {
          titleFormat: " MMMM D YYYY"
        },
        day: {
          titleFormat: 'D MMM, YYYY'
        }
      },

      select: function (start, end) {

        // on select we show the Sweet Alert modal with an input
        swal({
          title: 'Create an Event',
          html: '<div class="form-group">' +
            '<input class="form-control" placeholder="Event Title" id="input-field">' +
            '</div>',
          showCancelButton: true,
          confirmButtonClass: 'btn btn-success',
          cancelButtonClass: 'btn btn-danger',
          buttonsStyling: false
        }).then(function (result) {

          var eventData;
          event_title = $('#input-field').val();

          if (event_title) {
            eventData = {
              title: event_title,
              start: start,
              end: end
            };
            $calendar.fullCalendar('renderEvent', eventData, true); // stick? = true
          }

          $calendar.fullCalendar('unselect');

        });
      },
      editable: true,
      eventLimit: true, // allow "more" link when too many events


      // color classes: [ event-blue | event-azure | event-green | event-orange | event-red ]
      events: [
        {
          title: 'All Day Event',
          start: new Date(y, m, 1),
          className: 'event-default'
        },
        {
          id: 999,
          title: 'Repeating Event',
          start: new Date(y, m, d - 4, 6, 0),
          allDay: false,
          className: 'event-rose'
        },
        {
          id: 999,
          title: 'Repeating Event',
          start: new Date(y, m, d + 3, 6, 0),
          allDay: false,
          className: 'event-rose'
        },
        {
          title: 'Meeting',
          start: new Date(y, m, d - 1, 10, 30),
          allDay: false,
          className: 'event-green'
        },
        {
          title: 'Lunch',
          start: new Date(y, m, d + 7, 12, 0),
          end: new Date(y, m, d + 7, 14, 0),
          allDay: false,
          className: 'event-red'
        },
        {
          title: 'Md-pro Launch',
          start: new Date(y, m, d - 2, 12, 0),
          allDay: true,
          className: 'event-azure'
        },
        {
          title: 'Birthday Party',
          start: new Date(y, m, d + 1, 19, 0),
          end: new Date(y, m, d + 1, 22, 30),
          allDay: false,
          className: 'event-azure'
        },
        {
          title: 'Click for Creative Tim',
          start: new Date(y, m, 21),
          end: new Date(y, m, 22),
          url: 'http://www.creative-tim.com/',
          className: 'event-orange'
        },
        {
          title: 'Click for Google',
          start: new Date(y, m, 21),
          end: new Date(y, m, 22),
          url: 'http://www.creative-tim.com/',
          className: 'event-orange'
        }
      ]
    });
  },

  showNotification: function (from, align) {
    color = Math.floor((Math.random() * 4) + 1);

    $.notify({
      icon: "ti-gift",
      message: "Welcome to <b>Paper Dashboard</b> - a beautiful dashboard for every web developer."

    }, {
        type: type[color],
        timer: 4000,
        placement: {
          from: from,
          align: align
        }
      });
  },

  initDocumentationCharts: function () {
    //     	init single simple line chart
    var dataPerformance = {
      labels: ['6pm', '9pm', '11pm', '2am', '4am', '8am', '2pm', '5pm', '8pm', '11pm', '4am'],
      series: [
        [1, 6, 8, 7, 4, 7, 8, 12, 16, 17, 14, 13]
      ]
    };

    var optionsPerformance = {
      showPoint: false,
      lineSmooth: true,
      height: "200px",
      axisX: {
        showGrid: false,
        showLabel: true
      },
      axisY: {
        offset: 40,
      },
      low: 0,
      high: 16,
      height: "250px"
    };

    Chartist.Line('#chartPerformance', dataPerformance, optionsPerformance);

    //     init single line with points chart
    var dataStock = {
      labels: ['\'07', '\'08', '\'09', '\'10', '\'11', '\'12', '\'13', '\'14', '\'15'],
      series: [
        [22.20, 34.90, 42.28, 51.93, 62.21, 80.23, 62.21, 82.12, 102.50, 107.23]
      ]
    };

    var optionsStock = {
      lineSmooth: false,
      height: "200px",
      axisY: {
        offset: 40,
        labelInterpolationFnc: function (value) {
          return '$' + value;
        }

      },
      low: 10,
      height: "250px",
      high: 110,
      classNames: {
        point: 'ct-point ct-green',
        line: 'ct-line ct-green'
      }
    };

    Chartist.Line('#chartStock', dataStock, optionsStock);

    //      init multiple lines chart
    var dataSales = {
      labels: ['9:00AM', '12:00AM', '3:00PM', '6:00PM', '9:00PM', '12:00PM', '3:00AM', '6:00AM'],
      series: [
        [287, 385, 490, 562, 594, 626, 698, 895, 952],
        [67, 152, 193, 240, 387, 435, 535, 642, 744],
        [23, 113, 67, 108, 190, 239, 307, 410, 410]
      ]
    };

    var optionsSales = {
      lineSmooth: false,
      low: 0,
      high: 1000,
      showArea: true,
      height: "245px",
      axisX: {
        showGrid: false,
      },
      lineSmooth: Chartist.Interpolation.simple({
        divisor: 3
      }),
      showLine: true,
      showPoint: false,
    };

    var responsiveSales = [
      ['screen and (max-width: 640px)', {
        axisX: {
          labelInterpolationFnc: function (value) {
            return value[0];
          }
        }
      }]
    ];

    Chartist.Line('#chartHours', dataSales, optionsSales, responsiveSales);

    //      pie chart
    Chartist.Pie('#chartPreferences', {
      labels: ['62%', '32%', '6%'],
      series: [62, 32, 6]
    });

    //      bar chart
    var dataViews = {
      labels: ['Jan', 'Feb', 'Mar', 'Apr', 'Mai', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
      series: [
        [542, 443, 320, 780, 553, 453, 326, 434, 568, 610, 756, 895]
      ]
    };

    var optionsViews = {
      seriesBarDistance: 10,
      classNames: {
        bar: 'ct-bar'
      },
      axisX: {
        showGrid: false,

      },
      height: "250px"

    };

    var responsiveOptionsViews = [
      ['screen and (max-width: 640px)', {
        seriesBarDistance: 5,
        axisX: {
          labelInterpolationFnc: function (value) {
            return value[0];
          }
        }
      }]
    ];

    Chartist.Bar('#chartViews', dataViews, optionsViews, responsiveOptionsViews);

    //     multiple bars chart
    var data = {
      labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
      series: [
        [542, 543, 520, 680, 653, 753, 326, 434, 568, 610, 756, 895]
      ]
    };

    var options = {
      seriesBarDistance: 10,
      axisX: {
        showGrid: false
      },
      height: "245px"
    };

    var responsiveOptions = [
      ['screen and (max-width: 640px)', {
        seriesBarDistance: 5,
        axisX: {
          labelInterpolationFnc: function (value) {
            return value[0];
          }
        }
      }]
    ];

    Chartist.Line('#chartActivity', data, options, responsiveOptions);

  },


  registerFormValiaiton: function () {

    $('#registerFormValidation').validate({
      rules: {
        username: "required",
        email: {
          required: true,
          email: true,
        },
        password: {
          required: true,
          minlength: 5,
        },
        confirm_password: {
          required: true,
          minlength: 5,
          equalTo: "#password"
        },
        forename: "required",
        surname: "required",
        company_name: "required",
        search_address: "required",
        street_number: {
          required: true,
          number: true,
        },
        route: "required",
        locality: "required",
        postal_code: "required",
        country: "required",
        work_telephone: {
          required: true,
          number: true,
        },
        mobile_telephone: {
          required: true,
          number: true,
        },
      },
      messages: {
        username: "Pleeae enter your username",
        email1: "Please enter a valid email address",

        password: {
          required: "Please provide a password",
          minlength: "Your password must be at least 5 characters long"
        },
        confirm_password: {
          required: "Please provide a password",
          minlength: "Your password must be at least 5 characters long",
          equalTo: "Please enter the same password as above"
        },
        forename: "Please enter your forename",
        surname: "Please enter your surname",
        company_name: "Please enter your company name",
        search_address: "Please enter your address",
        street_number: {
          required: "Please enter your address number",
          number: "The address number is not a number",
        },
        route: "Please enter yout route",
        locality: "Please enter you city",
        postal_code: "Please enter your postal code",
        country: "Please enter your country",
        work_telephone: {
          required: "Please enter your phone number",
          number: "The phone number is not a number",
        },
        mobile_telephone: {
          required: "Please enter your mobile number",
          number: "The mobile number is not a number",
        }
      }
    });
  },
  showSwalShops: function (type) {
    if (type == 'warning-message-and-cancel') {
      swal({
        title: 'Are you sure?',
        text: 'You will not be able to recover this imaginary file!',
        type: 'warning',
        showCancelButton: true,
        confirmButtonText: 'Yes, delete it!',
        cancelButtonText: 'No, keep it',
        confirmButtonClass: "btn btn-success btn-fill",
        cancelButtonClass: "btn btn-danger btn-fill",
        buttonsStyling: false
      }).then(function () {
        swal({
          title: 'Deleted!',
          text: 'Your imaginary file has been deleted.',
          type: 'success',
          confirmButtonClass: "btn btn-success btn-fill",
          buttonsStyling: false
        })
      }, function (dismiss) {
        // dismiss can be 'overlay', 'cancel', 'close', 'esc', 'timer'
        if (dismiss === 'cancel') {
          swal({
            title: 'Cancelled',
            text: 'Your imaginary file is safe :)',
            type: 'error',
            confirmButtonClass: "btn btn-info btn-fill",
            buttonsStyling: false
          })
        }
      })

    }
  },

  loginFormValidation: function () {
    $('#loginFormValidation').validate({
      rules: {
        username: "required",
        password: {
          required: true,
          minlength: 5,
        },
        role: "required",

      },
      messages: {
        username: "Pleeae ender your VAT number or username",

        password: {
          required: "Please provide a password",
          minlength: "Your password must be at least 5 characters long",
        },
        role: "Please select role"
      }
    });
  },


  createNewHairdresserFormSubmit: function () {
    $("#hairdresser_name_exists").hide();
    $('#addHairdresser').on('submit', function (event) {
      event.preventDefault();
      if ($.trim($("#hairdresser_name").val()) === "" || $.trim($("#hairdresser_mobile_phone").val()) === "" || $.trim($("#selectstore").val()) === "" || $.trim($("#hairdresser_mobile_phone").val()) === "" || $.trim($("#hairdresser_phone").val()) === "") {
        alert('you did not fill out one of the fields');
        return false;
      } else {
        $('#addhairdressersubmit').addClass('loaderform');
        var user_is = $('input[name= UserID]').val();
        var usr_id = parseInt(user_is);



        var stores = $('#selectstore').val();



        var hairdresser_name = $('input[name=hairdresser_name]').val();
        var hairdresser_mobile_phone = $('input[name=hairdresser_mobile_phone]').val();
        var hairdresser_phone = $('input[name=hairdresser_phone]').val();

        var display_order = $('input[name=display_order]').val();
        display_order = parseFloat(display_order);


        var formData = {
          "user_id": usr_id,
          "stores": stores,
          "hairdresser_name": hairdresser_name,
          "hairdresser_mobile_phone": hairdresser_mobile_phone,
          "hairdresser_phone": hairdresser_phone,
          "display_order": display_order,
        };
        $.ajax({
          type: 'post',
          url: '/admin/api/create-hairdresser',
          data: JSON.stringify(formData),
          dataType: 'json',
          contentType: "application/json",

          success: function (data) {
            var status = data.status;
            console.log(status);
            if (status == 501) {
              $("#hairdresser_name_exists").show();
              $("#hairdresser_name_exists").text(data.description);
              $('#addhairdressersubmit').removeClass('loaderform');
              return;
            } else {
              $('#productsModalHorizontal').modal('toggle');
              $('#addhairdressersubmit').removeClass('loaderform');
              swal({
                title: "Good job!",
                text: data.description,
                buttonsStyling: false,
                confirmButtonClass: "btn btn-success btn-fill",
                type: "success"
              }).then(function () {
                window.location.reload();
              });
            }
          }
        });
      }
    });
  },




  createNewShopFormValidationSubmit: function () {
    $("#report").hide();
    $('#createNewShopFormValidation').on('submit', function (event) {
      event.preventDefault();
      if ($.trim($("#vat_number").val()) === "" || $.trim($("#company_name").val()) === "" || $.trim($("#work_telephone").val()) === "") {
        alert('you did not fill out one of the fields In New Shop Form');
        return false;
      } else {
        var user_is = $('input[name= UserId]').val();
        var usr_id = parseInt(user_is);
        var shopLogoImageFile = $('input[name=shopLogoImageFile]');
        var street_number = $("input[name=street_number]").val();
        street_number = parseInt(street_number);
        var formData = {
          'user_id': usr_id,
          'company_name': $("input[name=company_name]").val(),
          'vat_number': $("input[name=vat_number]").val(),
          'tax_office': $("input[name=tax_office]").val(),
          'work_telephone': $("input[name=work_telephone]").val(),
          'mobile_telephone': $("input[name=mobile_telephone]").val(),

          'company_street_number': street_number,
          'company_address': $("input[name=route]").val(),
          'company_city': $("input[name=locality]").val(),
          'company_state': $("input[name=administrative_area_level_1]").val(),
          'company_zip_code': $("input[name=postal_code]").val(),
          'company_country': $("input[name=country]").val(),
          'password': $("input[name=password]").val(),

          'include_bank_holidays': $('input[name="include_bank_holidays"]:checked').val(),
          'bank_holidays_country': $("#select_coundry option:selected").val()
        };
        //  debugger;
        //  return false;
        // var shopJSON = JSON.stringify(formData);

        // var fd = new FormData();
        //  fd.append('shopJSON', shopJSON);
        // if (shopLogoImageFile.length && shopLogoImageFile[0].files.length) {
        //   fd.append('shopLogoImageFile', shopLogoImageFile[0].files[0]);
        // }

        $.ajax({
          type: 'post',
          url: '/admin/api/create-shop',
          data: JSON.stringify(formData),
          dataType: 'json',
          contentType: "application/json",

          success: function (data) {
            var status = data.status;
            console.log(status);
            if (status == 501) {
              $("#report").text(data.description);
              $("#report").show();
              return;
            } else {
              $("#report").hide();
              $('#myModalHorizontal').modal('toggle');
              $.notify({
                icon: 'ti-user',
                message: data.description,



              }, {
                  placement: {
                    from: "top",
                    align: "center"
                  },
                  type: 'success',
                  timer: 2000,
                });
              setTimeout(function () {
                window.location.reload();
              }, 2000);

            }



          }
        });
      }
    });
  },
  updateShopFormValidationSubmit: function () {
    $('#editShopFormValidation').on('submit', function (event) {
      event.preventDefault();

      if ($.trim($("input[name=evat_number]").val()) === "" || $.trim($("input[name=ecompany_name]").val()) === "") {
        alert('you did not fill out one of the fields');
        return false;
      } else {
        var user_is = $('input[name= UserID]').val();
        var usr_id = parseInt(user_is);
        var id = $('input[name= Id]').val();
        var id = parseInt(id);
        var shopLogoImageFile = $('input[name=eshopLogoImageFile]');
        var street_number = $("input[name=estreet_number]").val();
        street_number = parseInt(street_number);
        var formData = {
          'user_id': usr_id,
          'company_name': $("input[name=ecompany_name]").val(),
          'vat_number': $("input[name=evat_number]").val(),
          'tax_office': $("input[name=etax_office]").val(),
          'work_telephone': $("input[name=ework_telephone]").val(),
          'mobile_telephone': $("input[name=emobile_telephone]").val(),

          'company_street_number': street_number,
          'company_address': $("input[name=eroute]").val(),
          'company_city': $("input[name=elocality]").val(),
          'company_state': $("input[name=eadministrative_area_level_1]").val(),
          'company_zip_code': $("input[name=epostal_code]").val(),
          'company_country': $("input[name=ecountry]").val(),
          'bank_holidays_country': $("#eselect_coundry option:selected").val()
        };

        // var shopJSON = JSON.stringify(formData);

        // var fd = new FormData();
        //  fd.append('shopJSON', shopJSON);
        // if (shopLogoImageFile.length && shopLogoImageFile[0].files.length) {
        //   fd.append('shopLogoImageFile', shopLogoImageFile[0].files[0]);
        // }

        $.ajax({
          type: 'post',
          url: '/admin/api/update-single-shop/' + id,
          data: JSON.stringify(formData),
          dataType: 'json',
          contentType: "application/json",

          success: function (data) {
            var status = data.status;
            console.log(status);
            if (status == 500) {
              $("#vat-exists").addClass("show");
              return;
            } else {
              $('#editModalHorizontal').modal('toggle');
              $.notify({
                icon: 'ti-user',
                message: data.description,



              }, {
                  placement: {
                    from: "top",
                    align: "center"
                  },
                  type: 'success',
                  timer: 2000,
                });
              setTimeout(function () {
                window.location.reload();
              }, 2000);

            }



          }
        });
      }
    });
  },


  enable_disable_subcategory: function (checked, id) {
    console.log(checked, id);
    id = parseInt(id);
    // var id = parseInt($(this).attr('data-id'));

    var formData = {
      'is_active': checked,
    }

    $.ajax({
      type: 'post',
      url: '/admin/api/enable-disable-sub-category/' + id,
      data: JSON.stringify(formData),

      success: function (data) {
        demo.showSwal('update-message');

      },


      contentType: 'application/json',
      dataType: 'json'
    });



  },
  enable_disable_category: function (checked, id) {
    console.log(checked, id);
    id = parseInt(id);



    var formData = {
      'is_active': checked,
    }

    $.ajax({
      type: 'post',
      url: '/admin/api/enable-disable-category/' + id,
      data: JSON.stringify(formData),

      success: function (data) {
        demo.showSwal('update-message');

      },


      contentType: 'application/json',
      dataType: 'json'
    });

  },
  enable_disable_service: function (checked, id, store_id) {
    id = parseInt(id);
    store_id = parseInt(store_id);
    var formData = {
      'is_active': checked,
    }

    $.ajax({
      type: 'post',
      url: '/admin/api/enable-disable-service/' + store_id + "/" + id,
      data: JSON.stringify(formData),

      success: function (data) {
        demo.showSwal('update-message');

      },
      contentType: 'application/json',
      dataType: 'json'
    });
  },
  has_formula: function (checked, id, store_id) {
    store_id = parseInt(store_id);

    var formData = {
      'is_active': checked,
    }


    $.ajax({
      type: 'post',
      url: '/admin/api/enable-disable-has-formula/' + store_id + "/" + id,
      data: JSON.stringify(formData),

      success: function (data) {
        demo.showSwal('update-message');

      },
      contentType: 'application/json',
      dataType: 'json'
    });
  },
  enable_disable_hairdresser: function (checked, id, store_id) {
    id = parseInt(id);
    store_id = parseInt(store_id);
    var formData = {
      'is_active': checked,
    }

    $.ajax({
      type: 'post',
      url: '/admin/api/enable-disable-hairdresser/' + store_id + "/" + id,
      data: JSON.stringify(formData),

      success: function (data) {
        swal({
          title: "Good job!",
          text: "Successfully updated hairdresser to the database",
          buttonsStyling: false,
          confirmButtonClass: "btn btn-success btn-fill",
          type: "success"
        }).then(function () {
          window.location.reload();
        });

      },


      contentType: 'application/json',
      dataType: 'json'
    });


  }
},



  $(document).on('click', '#close-preview', function () {
    $('.image-preview').popover('hide');
    // Hover befor close the preview
    $('.image-preview').hover(
      function () {
        $('.image-preview').popover('show');
      },
      function () {
        $('.image-preview').popover('hide');
      }
    );
  });

$(function () {
  // Create the close button
  var closebtn = $('<button/>', {
    type: "button",
    text: 'x',
    id: 'close-preview',
    style: 'font-size: initial;',
  });
  closebtn.attr("class", "close pull-right");
  // Set the popover default content
  $('.image-preview').popover({
    trigger: 'manual',
    html: true,
    title: "<strong>Preview</strong>" + $(closebtn)[0].outerHTML,
    content: "There's no image",
    placement: 'bottom'
  });
  // Clear event
  $('.image-preview-clear').click(function () {
    $('.image-preview').attr("data-content", "").popover('hide');
    $('.image-preview-filename').val("");
    $('.image-preview-clear').hide();
    $('.image-preview-input input:file').val("");
    $(".image-preview-input-title").text("Browse");
  });
  // Create the preview image
  $(".image-preview-input input:file").change(function () {
    var img = $('<img/>', {
      id: 'dynamic',
      width: 250,
      height: 200
    });
    var file = this.files[0];
    var reader = new FileReader();
    // Set preview image into the popover data-content
    reader.onload = function (e) {
      $(".image-preview-input-title").text("Change");
      $(".image-preview-clear").show();
      $(".image-preview-filename").val(file.name);
      img.attr('src', e.target.result);
      $(".image-preview").attr("data-content", $(img)[0].outerHTML).popover("show");
    }
    reader.readAsDataURL(file);
  });
});







