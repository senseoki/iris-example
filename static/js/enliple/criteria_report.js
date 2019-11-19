const $customerID = $('#customer-id');
    const $excel = $('#excel');

$(document).ready(function() {
    

    init();
    event();
});

function init() {
    //Initialize Select2 Elements
    $('.select2bs4').select2({
        theme: 'bootstrap4'
    });
}

function event() {
    $excel.click((e) => {
        e.preventDefault();

        axios({
            method: 'GET',
            url: 'http://localhost:8080/ads/downReport/report.csv',                 
            responseType: 'blob'
        })    
        .then(response =>{        
            const url = window.URL.createObjectURL(new Blob([response.data], { type: response.headers['content-type'] }));
            const link = document.createElement('a');
            link.href = url;
            link.setAttribute('download', 'test_'+moment(moment.now()).format("DD_MM_YYYY_HH_mm_ss")+'.csv');
            document.body.appendChild(link);
            link.click();

            alert(1)
        });
    });
    // $excel.click((e) => {
    //     e.preventDefault();
    //     $.ajax({
    //         url: "http://localhost:8080/ads/downReport/report.csv",
    //         type: 'GET',
    //         success: function (res) {
    //             // console.log(res)
    //             alert(1)
    //         },
    //         error: function (err) {
                
    //         }
    //     });
    // });
}