<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>{{.Title}}</title>
    </head>
    <body onload = "JavaScript:AutoRefresh(5000);">
       <table border="0" width="50%" align="center" style="font-family: 'Open Sans', Arial, sans-serif;">
           <tr bgcolor="#fdcb6e">
               <th style="padding: 5px 5px 5px 5px;">
                   FACTOR
               </th>
               <th style="padding: 5px 5px 5px 5px;">
                   VALUE
               </th>
               <th style="padding: 5px 5px 5px 5px;">
                   STATUS
               </th>
            </tr>
            <tr bgcolor="#ecf0f1">
                <td style="padding: 5px 5px 5px 5px;" align="center">
                    WIND
                </td>
                <td style="padding: 5px 5px 5px 5px;" align="center">
                    {{.NlAngin}}
                </td>
                <td style="padding: 5px 5px 5px 5px;" align="center">
                    {{.StAngin}}
                </td>
            </tr>
            <tr bgcolor="#bdc3c7">
                <td style="padding: 5px 5px 5px 5px;" align="center">
                    WATER
                </td>
                <td style="padding: 5px 5px 5px 5px;" align="center">
                    {{.NlAir}}
                </td>
                <td style="padding: 5px 5px 5px 5px;" align="center">
                    {{.StAir}}
                </td>
            </tr>
       </table>
    </body>
</html>
<script type = "text/JavaScript">
 <!--
    function AutoRefresh( t ) {
       setTimeout("location.reload(true);", t);
    }
 //-->
</script>