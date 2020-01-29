producthelpers = {

  addProductsValidation: function () {
    $(document).ready(function () {
      $("#product_name_exists").hide();

      var user_is = $('input[name=EditUserID]').val();
      var usr_id = parseInt(user_is);

      var url = '/admin/api/all-stores-per-user/' + usr_id;
      $.getJSON(url, function (data) {
        for (let i = 0; i < data["data"].length; i++) {
          $('#selectstore').append($('<option>').text(data["data"][i].company_name).attr('value', data["data"][i].id));

          console.log(data["data"][i].company_name);
        };
        $('.selectpicker').selectpicker('refresh');

      });

      $('#product_ml').on("input", function () {
        var product_ml = $('input[name=product_ml]').val();
        product_ml = parseInt(product_ml);

        var product_price = $('input[name=product_price]').val();
        product_price = parseFloat(product_price);

        var mlperprice = (product_price * 100) / product_ml;
        var intvalue = Math.ceil(mlperprice);
        $('input[name=product_ml_per_price').val(intvalue);
      });

      $('#eproduct_ml').on("input", function () {
        var eproduct_ml = $('input[name=eproduct_ml]').val();
        eproduct_ml = parseInt(eproduct_ml);

        var eproduct_price = $('input[name=eproduct_price]').val();
        eproduct_price = parseFloat(eproduct_price);

        var emlperprice = (eproduct_price * 100) / eproduct_ml;
        var eintvalue = Math.ceil(emlperprice);
        $('input[name=product_ml_per_price').val(eintvalue);
      });
    });

  }


}