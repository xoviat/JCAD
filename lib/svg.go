package lib

/*
   <style type="text/css">\n*[layerid="1"] {stroke:#FF0000;fill:#FF0000;}\n*[layerid="2"] {stroke:#0000FF;fill:#0000FF;}\n*[layerid="3"] {stroke:#FFCC00;fill:#FFCC00;}\n*[layerid="4"] {stroke:#66CC33;fill:#66CC33;}\n*[layerid="5"] {stroke:#808080;fill:#808080;}\n*[layerid="6"] {stroke:#800000;fill:#800000;}\n*[layerid="7"] {stroke:#800080;fill:#800080;stroke-opacity:0.7;fill-opacity:0.7;}\n*[layerid="8"] {stroke:#AA00FF;fill:#AA00FF;stroke-opacity:0.7;fill-opacity:0.7;}\n*[layerid="9"] {stroke:#6464FF;fill:#6464FF;}\n*[layerid="10"] {stroke:#FF00FF;fill:#FF00FF;}\n*[layerid="11"] {stroke:#C0C0C0;fill:#C0C0C0;}\n*[layerid="12"] {stroke:#FFFFFF;fill:#FFFFFF;}\n*[layerid="13"] {stroke:#33CC99;fill:#33CC99;}\n*[layerid="14"] {stroke:#5555FF;fill:#5555FF;}\n*[layerid="15"] {stroke:#F022F0;fill:#F022F0;}\n*[layerid="19"] {stroke:#66CCFF;fill:#66CCFF;}\n*[layerid="21"] {stroke:#999966;fill:#999966;}\n*[layerid="22"] {stroke:#008000;fill:#008000;}\n*[layerid="23"] {stroke:#00FF00;fill:#00FF00;}\n*[layerid="24"] {stroke:#BC8E00;fill:#BC8E00;}\n*[layerid="25"] {stroke:#70DBFA;fill:#70DBFA;}\n*[layerid="26"] {stroke:#00CC66;fill:#00CC66;}\n*[layerid="27"] {stroke:#9966FF;fill:#9966FF;}\n*[layerid="28"] {stroke:#800080;fill:#800080;}\n*[layerid="29"] {stroke:#008080;fill:#008080;}\n*[layerid="30"] {stroke:#15935F;fill:#15935F;}\n*[layerid="31"] {stroke:#000080;fill:#000080;}\n*[layerid="32"] {stroke:#00B400;fill:#00B400;}\n*[layerid="33"] {stroke:#2E4756;fill:#2E4756;}\n*[layerid="34"] {stroke:#99842F;fill:#99842F;}\n*[layerid="35"] {stroke:#FFFFAA;fill:#FFFFAA;}\n*[layerid="36"] {stroke:#99842F;fill:#99842F;}\n*[layerid="37"] {stroke:#2E4756;fill:#2E4756;}\n*[layerid="38"] {stroke:#3535FF;fill:#3535FF;}\n*[layerid="39"] {stroke:#8000BC;fill:#8000BC;}\n*[layerid="40"] {stroke:#43AE5F;fill:#43AE5F;}\n*[layerid="41"] {stroke:#C3ECCE;fill:#C3ECCE;}\n*[layerid="42"] {stroke:#728978;fill:#728978;}\n*[layerid="43"] {stroke:#39503F;fill:#39503F;}\n*[layerid="44"] {stroke:#0C715D;fill:#0C715D;}\n*[layerid="45"] {stroke:#5A8A80;fill:#5A8A80;}\n*[layerid="46"] {stroke:#2B937E;fill:#2B937E;}\n*[layerid="47"] {stroke:#23999D;fill:#23999D;}\n*[layerid="48"] {stroke:#45B4E3;fill:#45B4E3;}\n*[layerid="49"] {stroke:#215DA1;fill:#215DA1;}\n*[layerid="50"] {stroke:#4564D7;fill:#4564D7;}\n*[layerid="51"] {stroke:#6969E9;fill:#6969E9;}\n*[layerid="52"] {stroke:#9069E9;fill:#9069E9;}\n*[layerid="99"] {stroke:#00CCCC;fill:#00CCCC;}\n*[layerid="100"] {stroke:#CC9999;fill:#CC9999;}\n*[layerid="Hole"] {stroke:#222222;fill:#222222;}\n*[layerid="DRCError"] {stroke:#FAD609;fill:#FAD609;}\n*[fill="none"] {fill: none;}\n*[stroke="none"] {stroke: none;}\npath, polyline, polygon, line {stroke-linecap:round;}\ng[c_partid="part_pad"][layerid="1"] ellipse:not([c_padid]) {fill:#FF0000;}\ng[c_partid="part_pad"][layerid="1"]  polygon:not([c_padid]) {fill:#FF0000;}\ng[c_partid="part_pad"][layerid="1"]  polyline:not([c_padid]) {stroke:#FF0000;}\ng[c_partid="part_pad"][layerid="1"]  circle {fill:#FF0000;}\ng[c_partid="part_pad"][layerid="2"]  ellipse:not([c_padid]) {fill:#0000FF;}\ng[c_partid="part_pad"][layerid="2"]  polygon:not([c_padid]) {fill:#0000FF;}\ng[c_partid="part_pad"][layerid="2"]  polyline:not([c_padid]) {stroke:#0000FF;}\ng[c_partid="part_pad"][layerid="2"]  circle {fill:#0000FF;}\ng[c_partid="part_pad"][layerid="11"]  ellipse:not([c_padid]) {fill:#C0C0C0;}\ng[c_partid="part_pad"][layerid="11"]  polygon:not([c_padid]) {fill:#C0C0C0;}\ng[c_partid="part_pad"][layerid="11"]  polyline:not([c_padid]) {stroke:#C0C0C0;}\ng[c_partid="part_pad"][layerid="11"]  circle {fill:#C0C0C0;}\ng[c_partid="part_pad"][layerid] > circle[c_padhole] {fill: #222222;}\ng[c_partid="part_pad"][layerid] > polyline[c_padhole] {stroke:#222222;}\ng[c_partid="part_via"][layerid] > * + circle {fill: #222222;}\ng[c_partid="part_pad"] > polygon[c_padid] {stroke-linejoin: miter;stroke-miterlimit: 100;}\ng[c_partid="part_hole"] > circle {fill: #222222;}path, polyline, polygon {stroke-linejoin:round;}\nrect, circle, ellipse, polyline, line, polygon, path {shape-rendering:crispEdges;}\n</style>
   <rect x="3977.5" y="2986" width="33.8" height="28" fill="#000000" stroke="none" />
   <g c_origin="3994.437,3000" c_para="package`R0805`pre`R?`Contributor`layout`link`https://item.szlcsc.com/142685.html`3DModel`R0805_L2.0-W1.3-H0.6`" c_transformList="" xmlns="http://www.w3.org/2000/svg">
       <g c_partid="part_pad" c_etype="pinpart" c_origin="3998.37,3000" layerid="1" number="2" net="" plated="Y" c_shapetype="group" id="gge10" locked="0" c_rotation="0" c_width="3.937" c_height="6.1024" title="" c_shape="RECT" pasteExpansion="0" solderExpansion="0.4">
           <polygon points="3996.406 2996.949 4000.343 2996.949 4000.343 3003.051 3996.406 3003.051" layerid="7" stroke-width="0.8" c_padid="gge10" />
           <polygon points="3996.406 2996.949 4000.343 2996.949 4000.343 3003.051 3996.406 3003.051" layerid="5" stroke-width="0" c_padid="gge10" />
           <polygon points="3996.406 2996.949 4000.343 2996.949 4000.343 3003.051 3996.406 3003.051" stroke-width="0" />
           <circle c_padhole="1" cx="3998.374" cy="3000" r="0" stroke-width="0" />
       </g>
       <g c_partid="part_pad" c_etype="pinpart" c_origin="3990.5,3000" layerid="1" number="1" net="" plated="Y" c_shapetype="group" id="gge5" locked="0" c_rotation="0" c_width="3.937" c_height="6.1024" title="" c_shape="RECT" pasteExpansion="0" solderExpansion="0.4">
           <polygon points="3988.532 2996.949 3992.469 2996.949 3992.469 3003.051 3988.532 3003.051" layerid="7" stroke-width="0.8" c_padid="gge5" />
           <polygon points="3988.532 2996.949 3992.469 2996.949 3992.469 3003.051 3988.532 3003.051" layerid="5" stroke-width="0" c_padid="gge5" />
           <polygon points="3988.532 2996.949 3992.469 2996.949 3992.469 3003.051 3988.532 3003.051" stroke-width="0" />
           <circle c_padhole="1" cx="3990.5" cy="3000" r="0" stroke-width="0" />
       </g>
       <polyline points="3992.469 2995.984 3987.547 2995.984" stroke-width="1" c_shapetype="line" stroke-linecap="round" fill="none" layerid="3" net="" id="gge26" />
       <polyline points="3992.469 3004.016 3987.547 3004.016" stroke-width="1" c_shapetype="line" stroke-linecap="round" fill="none" layerid="3" net="" id="gge28" />
       <polyline points="3996.406 2995.984 4001.327 2995.984" stroke-width="1" c_shapetype="line" stroke-linecap="round" fill="none" layerid="3" net="" id="gge30" />
       <polyline points="3996.406 3004.016 4001.327 3004.016" stroke-width="1" c_shapetype="line" stroke-linecap="round" fill="none" layerid="3" net="" id="gge32" />
       <polyline points="3987.547 2995.984 3987.547 3004.016" stroke-width="1" c_shapetype="line" stroke-linecap="round" fill="none" layerid="3" net="" id="gge34" />
       <polyline points="4001.327 2995.984 4001.327 3004.016" stroke-width="1" c_shapetype="line" stroke-linecap="round" fill="none" layerid="3" net="" id="gge37" />
   </g>
*/

type EasyPackage struct {
	Title     string
	Pads      []*EasyPad
	Polylines []*EasyPolyline
}

type EasyPad struct{}

type EasyPolyline struct{}

type EasyCircle struct{}

type EasyPath struct{}
