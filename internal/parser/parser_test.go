package parser_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Toraja/todotxt-tui/internal/parser"
)

var _ = Describe("Parser", func() {
	var p parser.Parser

	BeforeEach(func() {
		p = parser.NewParser()
	})

	Describe("ParseLine", func() {
		Context("with simple task", func() {
			It("parses basic description", func() {
				task, err := p.ParseLine("Buy milk", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Description).To(Equal("Buy milk"))
				Expect(task.Completed).To(BeFalse())
				Expect(task.Priority).To(BeEmpty())
			})
		})

		Context("with priority", func() {
			It("parses priority A", func() {
				task, err := p.ParseLine("(A) Important task", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Priority).To(Equal("A"))
				Expect(task.Description).To(Equal("Important task"))
			})

			It("parses priority Z", func() {
				task, err := p.ParseLine("(Z) Low priority task", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Priority).To(Equal("Z"))
			})
		})

		Context("with creation date", func() {
			It("parses creation date", func() {
				task, err := p.ParseLine("2026-02-02 Buy groceries", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.CreationDate.Format("2006-01-02")).To(Equal("2026-02-02"))
				Expect(task.Description).To(Equal("Buy groceries"))
			})

			It("parses priority and creation date", func() {
				task, err := p.ParseLine("(A) 2026-02-02 Important task", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Priority).To(Equal("A"))
				Expect(task.CreationDate.Format("2006-01-02")).To(Equal("2026-02-02"))
				Expect(task.Description).To(Equal("Important task"))
			})
		})

		Context("with contexts", func() {
			It("extracts single context", func() {
				task, err := p.ParseLine("Buy milk @store", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Contexts).To(HaveKey("@store"))
			})

			It("extracts multiple contexts", func() {
				task, err := p.ParseLine("Task @work @urgent", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Contexts).To(HaveLen(2))
				Expect(task.Contexts).To(HaveKey("@work"))
				Expect(task.Contexts).To(HaveKey("@urgent"))
			})
		})

		Context("with projects", func() {
			It("extracts single project", func() {
				task, err := p.ParseLine("Buy groceries +household", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Projects).To(HaveKey("+household"))
			})

			It("extracts multiple projects", func() {
				task, err := p.ParseLine("Task +work +urgent", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Projects).To(HaveLen(2))
				Expect(task.Projects).To(HaveKey("+work"))
				Expect(task.Projects).To(HaveKey("+urgent"))
			})
		})

		Context("with metadata", func() {
			It("extracts key:value pairs", func() {
				task, err := p.ParseLine("Task due:2026-03-01", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Metadata).To(HaveKeyWithValue("due", "2026-03-01"))
			})

			It("extracts multiple metadata pairs", func() {
				task, err := p.ParseLine("Task due:2026-03-01 priority:high", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Metadata).To(HaveKeyWithValue("due", "2026-03-01"))
				Expect(task.Metadata).To(HaveKeyWithValue("priority", "high"))
			})
		})

		Context("with completed tasks", func() {
			It("parses completed marker", func() {
				task, err := p.ParseLine("x Buy milk", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Completed).To(BeTrue())
			})

			It("parses completion date", func() {
				task, err := p.ParseLine("x 2026-02-14 Buy milk", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Completed).To(BeTrue())
				Expect(task.CompletionDate.Format("2006-01-02")).To(Equal("2026-02-14"))
			})

			It("parses completion and creation dates", func() {
				task, err := p.ParseLine("x 2026-02-14 2026-02-02 Old task", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Completed).To(BeTrue())
				Expect(task.CompletionDate.Format("2006-01-02")).To(Equal("2026-02-14"))
				Expect(task.CreationDate.Format("2006-01-02")).To(Equal("2026-02-02"))
			})

			It("ignores priority for completed tasks", func() {
				task, err := p.ParseLine("x (A) 2026-02-14 Task", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Completed).To(BeTrue())
				Expect(task.Priority).To(BeEmpty())
			})
		})

		Context("with complex tasks", func() {
			It("parses full format task", func() {
				task, err := p.ParseLine("(A) 2026-02-02 Buy groceries @store +household due:2026-02-15", 1)
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Priority).To(Equal("A"))
				Expect(task.CreationDate.Format("2006-01-02")).To(Equal("2026-02-02"))
				Expect(task.Contexts).To(HaveKey("@store"))
				Expect(task.Projects).To(HaveKey("+household"))
				Expect(task.Metadata).To(HaveKeyWithValue("due", "2026-02-15"))
			})
		})

		Context("with invalid input", func() {
			It("returns error for empty line", func() {
				_, err := p.ParseLine("", 1)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("ParseFile", func() {
		Context("with multiple tasks", func() {
			It("parses all tasks", func() {
				content := `(A) 2026-02-02 Buy groceries @store
(B) Call mom @phone
x 2026-02-01 2026-01-31 Old completed task
Simple task without priority`

				reader := strings.NewReader(content)
				tasks, err := p.ParseFile(reader)

				Expect(err).ToNot(HaveOccurred())
				Expect(tasks).To(HaveLen(4))

				Expect(tasks[0].Priority).To(Equal("A"))
				Expect(tasks[0].Contexts).To(HaveKey("@store"))

				Expect(tasks[1].Priority).To(Equal("B"))
				Expect(tasks[1].Contexts).To(HaveKey("@phone"))

				Expect(tasks[2].Completed).To(BeTrue())

				Expect(tasks[3].Description).To(Equal("Simple task without priority"))
			})

			It("skips empty lines", func() {
				content := `Task 1

Task 2

Task 3`

				reader := strings.NewReader(content)
				tasks, err := p.ParseFile(reader)

				Expect(err).ToNot(HaveOccurred())
				Expect(tasks).To(HaveLen(3))
			})

			It("handles malformed lines gracefully", func() {
				content := `Valid task
x 2026-02-14 Another valid task`

				reader := strings.NewReader(content)
				tasks, err := p.ParseFile(reader)

				Expect(err).ToNot(HaveOccurred())
				Expect(tasks).To(HaveLen(2))
			})
		})

		Context("with empty file", func() {
			It("returns empty task list", func() {
				reader := strings.NewReader("")
				tasks, err := p.ParseFile(reader)

				Expect(err).ToNot(HaveOccurred())
				Expect(tasks).To(BeEmpty())
			})
		})
	})

	Describe("Serialize", func() {
		It("converts task to todo.txt format", func() {
			task := &parser.Task{
				Priority:    "A",
				Description: "Buy groceries",
			}

			result := p.Serialize(task)
			Expect(result).To(Equal("(A) Buy groceries"))
		})
	})

	Describe("Validate", func() {
		Context("with valid task", func() {
			It("accepts valid task", func() {
				task := &parser.Task{
					Priority:    "A",
					Description: "Buy milk @store +household",
					Contexts:    map[string]struct{}{"@store": {}},
					Projects:    map[string]struct{}{"+household": {}},
				}

				err := p.Validate(task)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("with invalid task", func() {
			It("rejects empty description", func() {
				task := &parser.Task{
					Description: "   ",
				}

				err := p.Validate(task)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("empty"))
			})

			It("rejects invalid priority", func() {
				task := &parser.Task{
					Priority:    "a",
					Description: "Task",
				}

				err := p.Validate(task)
				Expect(err).To(HaveOccurred())
			})

			It("rejects invalid context", func() {
				task := &parser.Task{
					Description: "Task",
					Contexts:    map[string]struct{}{"invalid": {}},
				}

				err := p.Validate(task)
				Expect(err).To(HaveOccurred())
			})

			It("rejects invalid project", func() {
				task := &parser.Task{
					Description: "Task",
					Projects:    map[string]struct{}{"invalid": {}},
				}

				err := p.Validate(task)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
