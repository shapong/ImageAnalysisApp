// $(document).ready(function() {
//     $.ajax({
//         // url: "https://api.havenondemand.com/1/api/sync/recognizeimages/v1&apikey=468f4fa3-530e-4898-8301-e98e32c43591"
//         //url: "http://rest-service.guides.spring.io/greeting"
//         url: "https://api.havenondemand.com/1/api/sync/ocrdocument/v1",
//         data: {
//             apikey: "468f4fa3-530e-4898-8301-e98e32c43591",
//             url: "https://www.havenondemand.com/sample-content/documents/hp_q1_2013.pdf"},
//         type: "POST",
//         success: function (data) {
//             //$('.greeting-content').append(data)
//             alert(JSON.stringify(data));
//         }
//     // }).then(function(data) {
//     //    $('.greeting-id').append(data.id);
//     //    $('.greeting-content').append(data);
//     // })
//     });
    
// });

function callHOD(api_key) {
    alert(api_key);
      $.ajax({
        // url: "https://api.havenondemand.com/1/api/sync/recognizeimages/v1&apikey=468f4fa3-530e-4898-8301-e98e32c43591"
        //url: "http://rest-service.guides.spring.io/greeting"
        url: "https://api.havenondemand.com/1/api/sync/ocrdocument/v1",
        data: {
            apikey: api_key,
            url: "https://www.havenondemand.com/sample-content/images/speccoll.jpg"},
        type: "POST",
        success: function (data) {
            //$('.greeting-content').append(data)
            //document.getElementById("log").innerText = JSON.stringify(data)
            alert(JSON.stringify(data));
            ()
        }
    // }).then(function(data) {
    //    $('.greeting-id').append(data.id);
    //    $('.greeting-content').append(data);
    // })
    });  
    
}