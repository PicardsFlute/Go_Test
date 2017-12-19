/**
 * Created by benjaminxerri on 12/10/17.
 */


$('#building-select').on('change',function(){
    $.ajax({
        url:"/admin/section/room/"+this.value,
        type:"GET",
        success:function(data){
            var jsonData = JSON.parse(data);
            var results = $('#room-results');
            //empty results
            results.empty();

            jsonData.forEach(function(value){
                var option = $("<option value="+value.RoomID +"> Number: " +value.RoomNumber + "-- Capacity:" + value.RoomCapacity + "</option>");

                // "---- <b>Type: </b>"+ value.RoomType +
                // option.add('value',value.RoomID);
                results.append(option);
            });
            console.log(jsonData)
        }
    })
});

