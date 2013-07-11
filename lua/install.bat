@rem Please set IUP_HOME first
@rem 
@echo Setting CGO flags ...

@set CGO_CFLAGS=-I%IUP_HOME%\include
@set CGO_LDFLAGS=-L%IUP_HOME%
@path=%IUP_HOME%;%path%

@echo Do buildings ...

@go install -x github.com/Archs/golua/lua

@echo Build Done

@pause
