## staticconfig on zoltar

The zoltar design is in `docs/design.md`. The staticconfig port can start once
its step 1 (the reshape) is done. Every staticconfig fixture must give the same
results before and after the port.

### Example: staticconfig's route loop

```go
p, err := z.Predict(
	ctx,
	appConfigure,
	zoltar.Unknown(appType, "the application"),
	configurer, // zoltar.Sentinel(configurerType)
)
if err != nil {
	return err
}

routes, err := p.ValueOf(ctx, routesCall.Call.Args[0], routesCall)
if err != nil {
	return err
}

for _, slice := range routes.Timelines() {
	elems, _ := slice.SliceElements() // ok: a single timeline
	speculative := !elems.IsComplete()

	for _, el := range elems.All() {
		_, indexKnown := el.Index()
		speculative := speculative || el.IsConditional() || !indexKnown

		for _, route := range el.Value().Timelines() {
			o, ok := route.Origin()
			if !ok {
				continue // partial: the route isn't known
			}
			via := o.Prediction()       // the dogma.ViaX() prediction
			pos := via.CallSite().Pos() // RoutedAt

			// The handler is ViaX's parameter, which relies only on Dogma's
			// public signature.
			fn := via.Function()
			h, err := via.ValueOf(ctx, fn.Params[0], fn.Blocks[0].Instrs[0])
			if err != nil {
				return err
			}
			hType, ok := h.DynamicType() // e.g. *app.Handler
			...
		}
	}
}
```

To check during the port: this relies on each route referring to a variable
allocated inside `ViaX`, so that it has an `Origin`. If Dogma's routes are
plain structs, staticconfig instead finds the `ViaX` calls in the function
itself, and reads the handler with `ValueOf` at each call.

### Example: should staticconfig follow this call?

staticconfig walks each argument itself, following its structure and using
`Origin` to stop at variables it has already visited:

```go
func mayHoldConfigurer(v zoltar.Value, seen map[zoltar.Origin]bool) bool {
	for _, t := range v.Timelines() {
		if o, ok := t.Origin(); ok {
			if seen[o] {
				continue
			}
			seen[o] = true
		}
		if eq, _ := t.Equal(configurer); eq {
			return true
		}
		if t.HasBlindSpots() && typeMayHold(t.StaticType(), configurerType) {
			return true // unknown, but could be the configurer
		}
		// Recurse into t.Elem(), t.Fields(), t.SliceElements(), t.MapEntries()
		// and, for a function value, Func.Captures().
	}
	return false
}
```

### Decided for staticconfig

- `dogma.ViaX(nil)`: zoltar sees that the call never returns. staticconfig
  reports "Configure() always panics" at the ViaX call, with a Related entry
  "nil handler passed to dogma.ViaX()". Recorded in the plan file.
- Recursion with differing arguments is cut off at the internal depth limit,
  not merged. Merging would change the memory model (one allocation standing
  for many variables) and is a separate research item.
