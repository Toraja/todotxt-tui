package parser_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Toraja/todotxt-tui/internal/parser"
)

func TestParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parser Suite")
}

var _ = Describe("Task", func() {
	var task *parser.Task

	BeforeEach(func() {
		task = &parser.Task{
			Priority:     "A",
			CreationDate: time.Date(2026, 2, 2, 0, 0, 0, 0, time.UTC),
			Description:  "Buy groceries @store +household",
			Contexts:     map[string]struct{}{"@store": {}},
			Projects:     map[string]struct{}{"+household": {}},
			Metadata:     make(map[string]string),
			LineNumber:   1,
		}
	})

	Describe("IsComplete", func() {
		Context("when task is not completed", func() {
			It("returns false", func() {
				Expect(task.IsComplete()).To(BeFalse())
			})
		})

		Context("when task is completed", func() {
			BeforeEach(func() {
				task.Completed = true
			})

			It("returns true", func() {
				Expect(task.IsComplete()).To(BeTrue())
			})
		})
	})

	Describe("SetPriority", func() {
		Context("with valid priority", func() {
			It("sets priority A", func() {
				err := task.SetPriority("A")
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Priority).To(Equal("A"))
				Expect(task.Modified).To(BeTrue())
			})

			It("sets priority Z", func() {
				err := task.SetPriority("Z")
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Priority).To(Equal("Z"))
			})

			It("clears priority with empty string", func() {
				err := task.SetPriority("")
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Priority).To(Equal(""))
				Expect(task.Modified).To(BeTrue())
			})
		})

		Context("with invalid priority", func() {
			It("rejects lowercase letter", func() {
				err := task.SetPriority("a")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("A-Z"))
			})

			It("rejects multiple characters", func() {
				err := task.SetPriority("AB")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("single letter"))
			})

			It("rejects numbers", func() {
				err := task.SetPriority("1")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("AddContext", func() {
		BeforeEach(func() {
			task.Contexts = make(map[string]struct{})
			task.Modified = false
		})

		Context("with valid context", func() {
			It("adds new context", func() {
				err := task.AddContext("@work")
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Contexts).To(HaveKey("@work"))
				Expect(task.Modified).To(BeTrue())
			})

			It("does not add duplicate context", func() {
				err := task.AddContext("@work")
				Expect(err).ToNot(HaveOccurred())
				err = task.AddContext("@work")
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Contexts).To(HaveLen(1))
			})
		})

		Context("with invalid context", func() {
			It("rejects context without @ prefix", func() {
				err := task.AddContext("work")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("@"))
			})
		})
	})

	Describe("RemoveContext", func() {
		BeforeEach(func() {
			task.Contexts = map[string]struct{}{"@work": {}, "@home": {}}
			task.Modified = false
		})

		It("removes existing context", func() {
			task.RemoveContext("@work")
			Expect(task.Contexts).ToNot(HaveKey("@work"))
			Expect(task.Contexts).To(HaveKey("@home"))
			Expect(task.Modified).To(BeTrue())
		})

		It("does nothing if context not present", func() {
			task.RemoveContext("@missing")
			Expect(task.Contexts).To(HaveLen(2))
			Expect(task.Modified).To(BeFalse())
		})
	})

	Describe("AddProject", func() {
		BeforeEach(func() {
			task.Projects = make(map[string]struct{})
			task.Modified = false
		})

		Context("with valid project", func() {
			It("adds new project", func() {
				err := task.AddProject("+personal")
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Projects).To(HaveKey("+personal"))
				Expect(task.Modified).To(BeTrue())
			})

			It("does not add duplicate project", func() {
				err := task.AddProject("+personal")
				Expect(err).ToNot(HaveOccurred())
				err = task.AddProject("+personal")
				Expect(err).ToNot(HaveOccurred())
				Expect(task.Projects).To(HaveLen(1))
			})
		})

		Context("with invalid project", func() {
			It("rejects project without + prefix", func() {
				err := task.AddProject("personal")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("+"))
			})
		})
	})

	Describe("RemoveProject", func() {
		BeforeEach(func() {
			task.Projects = map[string]struct{}{"+work": {}, "+personal": {}}
			task.Modified = false
		})

		It("removes existing project", func() {
			task.RemoveProject("+work")
			Expect(task.Projects).ToNot(HaveKey("+work"))
			Expect(task.Projects).To(HaveKey("+personal"))
			Expect(task.Modified).To(BeTrue())
		})

		It("does nothing if project not present", func() {
			task.RemoveProject("+missing")
			Expect(task.Projects).To(HaveLen(2))
			Expect(task.Modified).To(BeFalse())
		})
	})

	Describe("HasContext", func() {
		BeforeEach(func() {
			task.Contexts = map[string]struct{}{"@work": {}, "@home": {}}
		})

		It("returns true for existing context", func() {
			Expect(task.HasContext("@work")).To(BeTrue())
		})

		It("returns false for non-existing context", func() {
			Expect(task.HasContext("@missing")).To(BeFalse())
		})
	})

	Describe("HasProject", func() {
		BeforeEach(func() {
			task.Projects = map[string]struct{}{"+work": {}, "+personal": {}}
		})

		It("returns true for existing project", func() {
			Expect(task.HasProject("+work")).To(BeTrue())
		})

		It("returns false for non-existing project", func() {
			Expect(task.HasProject("+missing")).To(BeFalse())
		})
	})

	Describe("Complete", func() {
		BeforeEach(func() {
			task.Completed = false
			task.Modified = false
		})

		It("marks task as completed with given date", func() {
			completionDate := time.Date(2026, 2, 14, 0, 0, 0, 0, time.UTC)
			task.Complete(completionDate)
			Expect(task.Completed).To(BeTrue())
			Expect(task.CompletionDate).To(Equal(completionDate))
			Expect(task.Modified).To(BeTrue())
		})

		It("uses current date when given zero date", func() {
			task.Complete(time.Time{})
			Expect(task.Completed).To(BeTrue())
			Expect(task.CompletionDate).ToNot(BeZero())
		})
	})

	Describe("Uncomplete", func() {
		BeforeEach(func() {
			task.Completed = true
			task.CompletionDate = time.Date(2026, 2, 14, 0, 0, 0, 0, time.UTC)
			task.Modified = false
		})

		It("marks task as incomplete", func() {
			task.Uncomplete()
			Expect(task.Completed).To(BeFalse())
			Expect(task.CompletionDate).To(BeZero())
			Expect(task.Modified).To(BeTrue())
		})
	})

	Describe("String", func() {
		Context("with active task", func() {
			It("formats task with priority and creation date", func() {
				task.Priority = "A"
				task.CreationDate = time.Date(2026, 2, 2, 0, 0, 0, 0, time.UTC)
				task.Description = "Buy groceries @store +household"

				result := task.String()
				Expect(result).To(Equal("(A) 2026-02-02 Buy groceries @store +household"))
			})

			It("formats task without priority", func() {
				task.Priority = ""
				task.CreationDate = time.Date(2026, 2, 2, 0, 0, 0, 0, time.UTC)
				task.Description = "Buy milk"

				result := task.String()
				Expect(result).To(Equal("2026-02-02 Buy milk"))
			})

			It("formats task without creation date", func() {
				task.Priority = "B"
				task.CreationDate = time.Time{}
				task.Description = "Call mom"

				result := task.String()
				Expect(result).To(Equal("(B) Call mom"))
			})
		})

		Context("with completed task", func() {
			It("formats task with completion marker and date", func() {
				task.Completed = true
				task.CompletionDate = time.Date(2026, 2, 14, 0, 0, 0, 0, time.UTC)
				task.CreationDate = time.Date(2026, 2, 2, 0, 0, 0, 0, time.UTC)
				task.Description = "Old task"

				result := task.String()
				Expect(result).To(Equal("x 2026-02-14 2026-02-02 Old task"))
			})

			It("does not show priority for completed task", func() {
				task.Priority = "A"
				task.Completed = true
				task.CompletionDate = time.Date(2026, 2, 14, 0, 0, 0, 0, time.UTC)
				task.Description = "Done task"

				result := task.String()
				Expect(result).To(ContainSubstring("x 2026-02-14"))
				Expect(result).ToNot(ContainSubstring("(A)"))
			})
		})
	})
})
