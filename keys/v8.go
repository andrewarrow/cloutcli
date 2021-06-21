package keys

import (
	"fmt"

	"rogchap.com/v8go"
)

func RunV8(r, s string) string {
	ctx, _ := v8go.NewContext()

	js := `
function constructLength(arr, len) {
  if (len < 0x80) {
    arr.push(len);
    return;
  }
  var octets = 1 + (Math.log(len) / Math.LN2 >>> 3);
  arr.push(octets | 0x80);
  while (--octets) {
    arr.push((len >>> (octets << 3)) & 0xff);
  }
  arr.push(len);
};

function rmPadding(buf) {
  var i = 0;
  var len = buf.length - 1;
  while (!buf[i] && !(buf[i + 1] & 0x80) && i < len) {
    i++;
  }
  if (i === 0) {
    return buf;
  }
  return buf.slice(i);
};

function toDER(r, s) {
  if (r[0] & 0x80)
    r = [ 0 ].concat(r);
  if (s[0] & 0x80)
    s = [ 0 ].concat(s);

  r = rmPadding(r);
  s = rmPadding(s);

  while (!s[0] && !(s[1] & 0x80)) {
    s = s.slice(1);
  }
  var arr = [ 0x02 ];
  constructLength(arr, r.length);
  arr = arr.concat(r);
  arr.push(0x02);
  constructLength(arr, s.length);
  var backHalf = arr.concat(s);
  var res = [ 0x30 ];
  constructLength(res, backHalf.length);
  res = res.concat(backHalf);
	return res;
}

var r = [ %s ];
var s = [ %s ];
var result = toDER(r, s);
`

	ctx.RunScript(fmt.Sprintf(js, r, s), "value.js")

	val, _ := ctx.RunScript("result", "value.js")
	return fmt.Sprintf("%v", val)
}
