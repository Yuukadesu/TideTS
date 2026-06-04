package repl

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/hanami/tidets/cli/format"
	"github.com/hanami/tidets/client/session"
)

// Options 交互式 SQL CLI 配置。
type Options struct {
	In     io.Reader
	Out    io.Writer
	Err    io.Writer
	Prompt string
}

// Run 在已 Open 的 Session 上运行 SQL REPL，直到 exit/quit 或 EOF。
func Run(ctx context.Context, sess session.Session, opts Options) error {
	in := opts.In
	if in == nil {
		return fmt.Errorf("repl: input reader is required")
	}
	out := opts.Out
	if out == nil {
		out = io.Discard
	}
	errOut := opts.Err
	if errOut == nil {
		errOut = out
	}
	prompt := opts.Prompt
	if prompt == "" {
		prompt = "tidets> "
	}

	reader, ok := in.(*bufio.Reader)
	if !ok {
		reader = bufio.NewReader(in)
	}

	printHelp(out)
	var buf strings.Builder

	for {
		fmt.Fprint(out, prompt)
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if buf.Len() == 0 {
					fmt.Fprintln(out, "bye")
					return nil
				}
				line = "\n"
			} else {
				return err
			}
		}

		trimmed := strings.TrimSpace(line)
		if buf.Len() == 0 {
			if trimmed == "" {
				continue
			}
			if isMetaCommand(trimmed) {
				if handleMeta(trimmed, out) {
					return nil
				}
				continue
			}
		}

		buf.WriteString(line)
		sql := strings.TrimSpace(buf.String())
		if !strings.HasSuffix(sql, ";") {
			buf.WriteString(" ")
			continue
		}
		sql = strings.TrimSuffix(sql, ";")
		sql = strings.TrimSpace(sql)
		buf.Reset()

		if sql == "" {
			continue
		}

		res, err := sess.ExecuteSQL(ctx, sql)
		if err != nil {
			fmt.Fprintf(errOut, "ERROR: %v\n", err)
			continue
		}
		fmt.Fprintln(out, format.SQLResult(res))
	}
}

func isMetaCommand(line string) bool {
	switch strings.ToLower(line) {
	case "exit", "quit", "\\q", "help", "?":
		return true
	default:
		return false
	}
}

func handleMeta(cmd string, out io.Writer) (exit bool) {
	switch strings.ToLower(cmd) {
	case "exit", "quit", "\\q":
		fmt.Fprintln(out, "bye")
		return true
	case "help", "?":
		printHelp(out)
	default:
		printHelp(out)
	}
	return false
}

func printHelp(out io.Writer) {
	fmt.Fprintln(out, "TideTS SQL CLI (Session + ExecuteSQL)")
	fmt.Fprintln(out, "Supported SQL:")
	fmt.Fprintln(out, "  INSERT INTO <device>(<measurement>) VALUES (<ts>, <val>)[, (<ts>, <val>)...];")
	fmt.Fprintln(out, "  SELECT <measurement> FROM <device> [WHERE time <op> <n> [AND ...]] [LIMIT <n>];")
	fmt.Fprintln(out, "  SELECT COUNT(<measurement>) FROM <device> [WHERE time <op> <n> [AND ...]];")
	fmt.Fprintln(out, "  DELETE FROM <device>(<measurement>) WHERE time <op> <n> [AND ...];")
	fmt.Fprintln(out, "  CREATE TIMESERIES <device>(<measurement>) WITH DATATYPE=<TYPE>;")
	fmt.Fprintln(out, "  SHOW DEVICES [<path>[.**]];")
	fmt.Fprintln(out, "  SHOW TIMESERIES <device>;")
	fmt.Fprintln(out, "Meta: help, exit, quit")
	fmt.Fprintln(out)
}
