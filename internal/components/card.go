package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// Card is a component that displays a card.
type Card struct {
	app.Compo
}

// newCard creates a new Card.
func newCard() *Card {
	return &Card{}
}

// Render renders the Card.
func (c *Card) Render() app.UI {
	return app.Div().Body(
		app.Div().ID("terminal"),
		app.Script().Src("/web/static/stack.js"),
		app.Script().Src("/web/static/ws-delegate.js"),
	)
}

func (c *Card) LoadScript() app.UI {
	return app.Raw(`
	<script>
	const xterm = new Terminal();
	xterm.open(document.getElementById("terminal"));

	const { master, slave } = openpty();

	termios = slave.ioctl("TCGETS");
	termios.iflag &= ~(ISTRIP | INLCR | IGNCR | ICRNL | IXON);
	termios.oflag &= ~(OPOST);
	termios.lflag &= ~(ECHO | ECHONL | ICANON | ISIG | IEXTEN);

	slave.ioctl("TCSETS", new Termios(termios.iflag, termios.oflag, termios.cflag, termios.lflag, termios.cc));
	xterm.loadAddon(master);
	const worker = new Worker("/web/static/worker.js"+location.search);

	var nwStack;
	var netParam = getNetParam();
	var workerImage = location.origin + "/web/static/out.wasm";
	if (netParam) {
		if (netParam.mode == 'delegate') {
			nwStack = delegate(worker, workerImage, netParam.param);
		} else if (netParam.mode == 'browser') {
			nwStack = newStack(worker, workerImage, new Worker("/web/static/stack-worker.js"+location.search), location.origin + "/c2w-net-proxy.wasm");
		}
	}
	if (!nwStack) {
		worker.postMessage({type: "init", imagename: workerImage});
	}
	new TtyServer(slave).start(worker, nwStack);

	function getNetParam() {
		var vars = location.search.substring(1).split('&');
		for (var i = 0; i < vars.length; i++) {
			var kv = vars[i].split('=');
			if (decodeURIComponent(kv[0]) == 'net') {
				return {
					mode: kv[1],
					param: kv[2],
				};
			}
		}
		return null;
	}
	</script>
	`)
}
