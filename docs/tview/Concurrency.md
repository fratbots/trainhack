Concurrency is a tricky subject in `tview`. In many places, the package gives you direct access to the data. For example, you can directly modify the fields of a [`TableCell`](https://godoc.org/github.com/rivo/tview#TableCell). (In retrospect, this may not have been a great decision but we want to remain backwards compatible so it is here to stay.) If you modify such an object in a goroutine while another goroutine calls [`Application.Draw()`](https://godoc.org/github.com/rivo/tview#Application.Draw) which reads from your object to refresh the screen, you will hit race conditions.

The obvious idea is to hide all fields and provide getters and setters which use mutexes to serialize access to the fields. However, in addition to adding a lot of boilerplate code, it would only serialize access locally. Functions like [`TreeNode.ExpandAll()`](https://godoc.org/github.com/rivo/tview#TreeNode.ExpandAll) which operate on an entire tree of objects would not be protected from data races. Note that the getters and setters of this package generally do not synchronize access to member variables.

To solve these issues, there are a couple of things that can be done:

## Event Handlers

Any event handlers you install, e.g. [`InputField.SetChangedFunc()`](https://godoc.org/github.com/rivo/tview#InputField.SetChangedFunc) or [`Table.SetSelectedFunc()`](https://godoc.org/github.com/rivo/tview#Table.SetSelectedFunc), are invoked from the main goroutine. It is safe to make changes to your primitives in these handlers. If they are invoked in response to a key event, you also don't need to call [`Application.Draw()`](https://godoc.org/github.com/rivo/tview#Application.Draw) as the application's main loop will do that for you.

## TextView Changes

The [`TextView`](https://godoc.org/github.com/rivo/tview#TextView) primitive implements the [`io.Writer`](https://golang.org/pkg/io/#Writer) interface. It is common to write to a `TextView` from a different goroutine. Therefore, `tview` provides [`TextView.SetChangedFunc()`](https://godoc.org/github.com/rivo/tview#TextView.SetChangedFunc) which notifies you when text was written to your `TextView`. However, contrary to the other handler functions, your handler will be invoked from the same goroutine that writes to the `TextView` so you need to take extra care in your handler.

It is always safe to call `Application.Draw()` from the handler, and most of the time, this is the only action needed in the handler anyway. [`TextView.HasFocus()`](https://godoc.org/github.com/rivo/tview#TextView.HasFocus) is also safe, in case you want to update only when the `TextView` has focus.

For all other actions on primitives, your handler should queue an update. The next section describes how this is done.

## Actions in Goroutines

If you make modifications to primitives from within a goroutine, to avoid race conditions, you will need to synchronize them with the main application loop. The functions that help you do this are [`Application.QueueUpdate()`](https://godoc.org/github.com/rivo/tview#Application.QueueUpdate) and [`Application.QueueUpdateDraw()`](https://godoc.org/github.com/rivo/tview#Application.QueueUpdateDraw). Here's an example:

```go
go func() {
  app.QueueUpdateDraw(func() {
    table.SetCellSimple(0, 0, "Foo bar")
  })
}()
```

`QueueUpdateDraw()` is like `QueueUpdate()` but it also calls [`Application.Draw()`](https://godoc.org/github.com/rivo/tview#Application.Draw) at the end. Depending on the granularity of your changes, you may not always want to redraw the screen. The availability of these two functions leaves the decision up to you.

It is also recommended to use `QueueUpdate()` if you perform a read-only operation on a primitive in a goroutine, unless you are absolutely certain that no other primitive or goroutine makes any changes to the primitive you're accessing.

Note that calling `Application.Draw()` is always safe. If all you need to do is refresh the screen, you don't need to wrap the call to `Draw()` in `QueueUpdate()`:

```go
go func() {
  app.Draw()
}()
```

See also the [[Timer]] page for a real-life example of how this is used.