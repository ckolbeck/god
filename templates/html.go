package templates
import "html/template"
var HTML = template.New("html")
func init() {
  template.Must(HTML.New("index.html").Parse("<html>\n	<head>\n		<title>\n			Go Database!\n		</title>\n		<link href=\"/css/{{.T}}/all.css\" rel=\"stylesheet\" media=\"screen\">\n		<script src=\"/js/{{.T}}/all.js\" type=\"text/javascript\"></script>\n	</head>\n	<body>		\n		<div id=\"chord_container\">\n			<canvas width=\"3000\" height=\"2000\" id=\"chord\"></canvas>\n		</div>\n		<div id=\"nodes_container\">\n			<table class=\"table table-striped\" id=\"nodes\">\n				<caption>nodes</caption>\n				<tr>\n					<th>address</th>\n					<th>position</th>\n				</tr>\n			</table>\n		</div>\n		<div id=\"node_container\">\n			<table class=\"table table-condensed\">\n				<caption>node</caption>\n				<tr>\n					<td>gob rpc address</td>\n					<td id=\"node_gob_addr\"></td>\n				</tr>\n				<tr>\n					<td>JSON/HTTP rpc address</td>\n					<td id=\"node_json_addr\"></td>\n				</tr>\n				<tr>\n					<td>position</td>\n					<td id=\"node_pos\"></td>\n				</tr>\n				<tr>\n					<td>owned keys</td>\n					<td id=\"node_owned_keys\"></td>\n				</tr>\n				<tr>\n					<td>held keys</td>\n					<td id=\"node_held_keys\"></td>\n				</tr>\n			</table>\n		</div>\n	</body>\n</html>\n"))
}